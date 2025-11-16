// Package db contains core data structures used in the application,
// including definitions of users, teams, and pull requests.
package db

import "time"

// PullRequest represents a pull request in the service.
type PullRequest struct {
	ID                string     `json:"pull_request_id"`
	Name              string     `json:"pull_request_name"`
	AuthorID          string     `json:"author_id"`
	Status            string     `json:"status"`
	AssignedReviewers []string   `json:"assigned_reviewers"`
	MergedAt          *time.Time `json:"merged_at"`
}

// NewPullRequest creates a new PullRequest instance.
func NewPullRequest(id string, name string, authorID string, status string, assignedReviewers []string) *PullRequest {
	return &PullRequest{
		ID:                id,
		Name:              name,
		AuthorID:          authorID,
		Status:            status,
		AssignedReviewers: assignedReviewers,
	}
}

// NewPullRequestMerged creates a new PullRequest instance with merged time.
func NewPullRequestMerged(id string, name string, authorID string, status string, assignedReviewers []string, mergedAt time.Time) *PullRequest {
	return &PullRequest{
		ID:                id,
		Name:              name,
		AuthorID:          authorID,
		Status:            status,
		AssignedReviewers: assignedReviewers,
		MergedAt:          &mergedAt,
	}
}
