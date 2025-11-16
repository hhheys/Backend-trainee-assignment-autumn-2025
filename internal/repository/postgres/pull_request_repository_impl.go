package postgres

import (
	db2 "AvitoPRService/internal/model/db"
	"database/sql"
	"errors"
	"math/rand"
	"time"
)

// PullRequestRepositoryImpl is the implementation of the PullRequestRepository interface.
type PullRequestRepositoryImpl struct {
	db             *sql.DB
	userRepository UserRepository
	teamRepository TeamRepository
}

// NewPullRequestRepositoryImpl creates a new PullRequestRepositoryImpl instance.
func NewPullRequestRepositoryImpl(db *sql.DB, userRepository UserRepository, teamRepository TeamRepository) *PullRequestRepositoryImpl {
	return &PullRequestRepositoryImpl{
		db:             db,
		userRepository: userRepository,
		teamRepository: teamRepository,
	}
}

// FindPullRequestByID finds a pull request by ID.
func (r *PullRequestRepositoryImpl) FindPullRequestByID(pullRequestID string) (*db2.PullRequest, error) {
	var pullRequest db2.PullRequest
	err := r.db.QueryRow("SELECT id, name, author_id, status, merged_at FROM pull_request WHERE id = $1", pullRequestID).Scan(
		&pullRequest.ID,
		&pullRequest.Name,
		&pullRequest.AuthorID,
		&pullRequest.Status,
		&pullRequest.MergedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPullRequestNotExists
		}
		return nil, err
	}
	var reviewers []string
	rows, err := r.db.Query("SELECT reviewer_id FROM pull_request_reviewer WHERE pull_request_id = $1", pullRequestID)
	if err != nil {
		return nil, err
	}
	var reviewerID string
	for rows.Next() {
		err := rows.Scan(&reviewerID)
		if err != nil {
			continue
		}
		reviewers = append(reviewers, reviewerID)
	}
	pullRequest.AssignedReviewers = reviewers
	return &pullRequest, nil
}

// CreatePullRequest creates a new pull request.
func (r *PullRequestRepositoryImpl) CreatePullRequest(pullRequestID string, pullRequestName string, authorID string) (*db2.PullRequest, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow("SELECT 1 FROM pull_request WHERE id = $1", pullRequestID).Scan()
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, ErrPullRequestAlreadyExists
	}

	author, err := r.userRepository.FindUserByID(authorID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	if !author.IsActive {
		return nil, ErrUserIsNotActive
	}

	if *author.TeamName == "" {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, ErrUserNoTeamFound
	}

	var pullRequest db2.PullRequest
	err = tx.QueryRow(
		"INSERT INTO pull_request(id, name, author_id) VALUES($1, $2, $3) RETURNING id, name, author_id, status",
		pullRequestID,
		pullRequestName,
		authorID,
	).Scan(
		&pullRequest.ID,
		&pullRequest.Name,
		&pullRequest.AuthorID,
		&pullRequest.Status,
	)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	team, err := r.teamRepository.FindTeamByName(*author.TeamName)
	if err != nil {
		if errors.Is(err, ErrTeamNotExists) {
			return nil, err
		}
		return nil, err
	}

	var reviewersCandidatesIDs []string
	for _, user := range team.Members {
		if user.ID != author.ID && user.IsActive {
			reviewersCandidatesIDs = append(reviewersCandidatesIDs, user.ID)
		}
	}
	pullRequest.AssignedReviewers = r.pickReviewers(reviewersCandidatesIDs)

	for _, reviewer := range pullRequest.AssignedReviewers {
		_, err := tx.Exec("INSERT INTO pull_request_reviewer(pull_request_id, reviewer_id) VALUES($1, $2)", pullRequestID, reviewer)
		if err != nil {
			continue
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &pullRequest, nil
}

// MergePullRequest merges a pull request.
func (r *PullRequestRepositoryImpl) MergePullRequest(pullRequestID string) (*db2.PullRequest, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	pr, err := r.FindPullRequestByID(pullRequestID)
	if err != nil {
		return nil, err
	}
	if pr.Status == "MERGED" {
		return pr, nil
	}
	err = tx.QueryRow("UPDATE pull_request SET status = 'MERGED', merged_at = NOW() WHERE id = $1 RETURNING status, merged_at", pullRequestID).Scan(
		&pr.Status,
		&pr.MergedAt,
	)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func (r *PullRequestRepositoryImpl) pickReviewers(reviewCandidates []string) []string {
	var reviewers []string
	if len(reviewCandidates) == 0 {
		reviewers = make([]string, 0)
		return reviewers
	} else if len(reviewCandidates) <= 2 {
		return reviewCandidates
	}
	rand.Shuffle(len(reviewers), func(i, j int) {
		reviewers[i], reviewers[j] = reviewers[j], reviewers[i]
	})
	return reviewers[:2]
}

// ReassignReviewer reassigns a reviewer to a pull request.
func (r *PullRequestRepositoryImpl) ReassignReviewer(pullRequestID, reviewerID string) (*db2.PullRequest, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	pr, err := r.FindPullRequestByID(pullRequestID)
	if err != nil {
		return nil, err
	}
	if pr.Status == "MERGED" {
		return pr, ErrPullRequestMergedReassign
	}

	if !contains(pr.AssignedReviewers, reviewerID) {
		return pr, ErrUserIsNotAssignedToPR
	}

	author, err := r.userRepository.FindUserByID(pr.AuthorID)
	if err != nil {
		return nil, err
	}

	// Get candidates
	rows, err := tx.Query(`
        SELECT id 
        FROM users 
        WHERE team_name = $1 
          AND is_active = TRUE 
          AND id NOT IN ($2, $3)
    `, author.TeamName, pr.AuthorID, reviewerID)
	if err != nil {
		return pr, ErrNoActiveReplacementCandidates
	}

	var candidates []string
	for rows.Next() {
		var id string
		if scanErr := rows.Scan(&id); scanErr == nil {
			candidates = append(candidates, id)
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	// Exclude already assigned reviewers
	for _, id := range pr.AssignedReviewers {
		candidates = excludeValueInSlice(candidates, id)
	}

	if len(candidates) == 0 {
		return pr, ErrNoActiveReplacementCandidates
	}

	// Delete old from DB
	if _, err = tx.Exec(`
        DELETE FROM pull_request_reviewer 
        WHERE pull_request_id = $1 AND reviewer_id = $2
    `, pr.ID, reviewerID); err != nil {
		return nil, err
	}

	// Insert new reviewer into DB
	newReviewerID := candidates[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(candidates))]
	if _, err = tx.Exec(`
        INSERT INTO pull_request_reviewer(pull_request_id, reviewer_id) 
        VALUES($1, $2)
    `, pullRequestID, newReviewerID); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	pr.AssignedReviewers = append(
		excludeValueInSlice(pr.AssignedReviewers, reviewerID),
		newReviewerID,
	)

	return pr, nil
}

// contains checks if a string is in a slice.
func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

// excludeValueInSlice excludes a value from a slice.
func excludeValueInSlice(slice []string, value string) []string {
	i := 0
	for _, v := range slice {
		if v != value {
			slice[i] = v
			i++
		}
	}
	return slice[:i]
}
