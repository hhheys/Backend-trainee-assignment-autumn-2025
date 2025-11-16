package postgres

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryImpl_SetIsActive(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewUserRepositoryImpl(db)

	userID := "123"
	isActive := true
	rows := sqlmock.NewRows([]string{"id", "username", "team_name", "is_active"}).
		AddRow(userID, "testuser", "teamA", isActive)

	mock.ExpectQuery("UPDATE users SET is_active = \\$1 WHERE id = \\$2 RETURNING id, username, team_name, is_active").
		WithArgs(isActive, userID).
		WillReturnRows(rows)

	user, err := repo.SetIsActive(userID, isActive)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "teamA", *user.TeamName)
	assert.Equal(t, isActive, user.IsActive)

	// Checking for all mocks to complete
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryImpl_SetIsActive_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewUserRepositoryImpl(db)
	userID := "notexist"
	isActive := true

	mock.ExpectQuery("UPDATE users SET is_active = \\$1 WHERE id = \\$2 RETURNING id, username, team_name, is_active").
		WithArgs(isActive, userID).
		WillReturnError(sql.ErrNoRows)

	user, err := repo.SetIsActive(userID, isActive)
	assert.Nil(t, user)
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestUserRepositoryImpl_FindUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewUserRepositoryImpl(db)
	userID := "123"
	isActive := true

	rows := sqlmock.NewRows([]string{"id", "username", "team_name", "is_active"}).
		AddRow(userID, "testuser", "teamA", isActive)

	mock.ExpectQuery("SELECT id, username, team_name, is_active FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	user, err := repo.FindUserByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "teamA", *user.TeamName)
	assert.Equal(t, isActive, user.IsActive)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryImpl_FindUserByID_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close() //nolint:errcheck

	repo := NewUserRepositoryImpl(db)
	userID := "notexist"

	mock.ExpectQuery("SELECT id, username, team_name, is_active FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	user, err := repo.FindUserByID(userID)
	assert.Nil(t, user)
	assert.ErrorIs(t, err, ErrUserNotFound)
}
