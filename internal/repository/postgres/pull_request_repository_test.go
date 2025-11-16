package postgres

import (
	db2 "AvitoPRService/internal/model/db"
	"AvitoPRService/internal/model/dto"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct{}

func (m *mockUserRepository) FindUserByID(userID string) (*db2.User, error) {
	if userID == "active_user" {
		teamName := "team1" // обязательно не nil
		return &db2.User{
			ID:       "active_user",
			IsActive: true,
			TeamName: &teamName,
		}, nil
	}
	return nil, errors.New("user not found")
}

func (m *mockUserRepository) SetIsActive(userID string, isActive bool) (*db2.User, error) {
	user, err := m.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	user.IsActive = isActive
	return user, nil
}

// --- Мок TeamRepository ---
type mockTeamRepository struct{}

func (m *mockTeamRepository) FindTeamByName(name string) (*db2.Team, error) {
	if name == "team1" {
		return &db2.Team{
			Name: "team1",
			Members: []db2.User{
				{ID: "user1", IsActive: true},
				{ID: "user2", IsActive: true},
			},
		}, nil
	}
	return nil, ErrTeamNotExists
}

// Новые заглушки
func (m *mockTeamRepository) CreateTeam(teamName string, members []dto.TeamMemberDto) (*db2.Team, error) {
	teamMembers := make([]db2.User, 0, len(members))
	for _, member := range members {
		teamMembers = append(teamMembers, db2.User{ID: member.ID, IsActive: true})
	}
	return db2.NewTeam(teamName, teamMembers), nil
}

func (m *mockTeamRepository) IsTeamExists(teamName string) (bool, error) {
	if teamName == "team1" {
		return true, nil
	}
	return false, nil
}

func TestFindPullRequestByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	userRepository := NewUserRepositoryImpl(db)
	teamRepository := NewTeamRepositoryImpl(db)

	repo := NewPullRequestRepositoryImpl(db, userRepository, teamRepository)

	// Мокирование основного запроса pull_request
	rows := sqlmock.NewRows([]string{"id", "name", "author_id", "status", "merged_at"}).
		AddRow("pr1", "PR 1", "user1", "OPEN", nil)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, author_id, status, merged_at FROM pull_request WHERE id = $1")).
		WithArgs("pr1").WillReturnRows(rows)

	// Мокирование запроса reviewers
	reviewerRows := sqlmock.NewRows([]string{"reviewer_id"}).AddRow("user2")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT reviewer_id FROM pull_request_reviewer WHERE pull_request_id = $1")).
		WithArgs("pr1").WillReturnRows(reviewerRows)

	pr, err := repo.FindPullRequestByID("pr1")
	assert.NoError(t, err)
	assert.Equal(t, "pr1", pr.ID)
	assert.Equal(t, []string{"user2"}, pr.AssignedReviewers)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreatePullRequest_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	userRepository := &mockUserRepository{}
	teamRepository := &mockTeamRepository{}

	repo := NewPullRequestRepositoryImpl(db, userRepository, teamRepository)

	mock.ExpectBegin()

	// Проверка существования pull request
	mock.ExpectQuery(regexp.QuoteMeta("SELECT 1 FROM pull_request WHERE id = $1")).
		WithArgs("pr1").
		WillReturnError(sql.ErrNoRows)

	// Вставка нового pull request
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO pull_request(id, name, author_id) VALUES($1, $2, $3) RETURNING id, name, author_id, status")).
		WithArgs("pr1", "PR 1", "active_user").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author_id", "status"}).
			AddRow("pr1", "PR 1", "active_user", "OPEN"))

	// Вставка assigned reviewers
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO pull_request_reviewer(pull_request_id, reviewer_id) VALUES($1, $2)")).
		WithArgs("pr1", "user2").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	pr, err := repo.CreatePullRequest("pr1", "PR 1", "active_user")
	assert.NoError(t, err)
	assert.Equal(t, "pr1", pr.ID)
	assert.Contains(t, pr.AssignedReviewers, "user2")
	assert.NoError(t, mock.ExpectationsWereMet())
}
