package postgres

import (
	db2 "AvitoPRService/internal/model/db"
	"database/sql"
	"errors"
)

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	SetIsActive(userID string, isActive bool) (*db2.User, error)
	FindUserByID(userID string) (*db2.User, error)
}

// UserRepositoryImpl implements the UserRepository interface.
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl.
func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// SetIsActive sets the is_active status of a user.
func (r *UserRepositoryImpl) SetIsActive(userID string, isActive bool) (*db2.User, error) {
	var user db2.User
	err := r.db.QueryRow("UPDATE users SET is_active = $1 WHERE id = $2 RETURNING id, username, team_name, is_active", isActive, userID).Scan(&user.ID, &user.Username, &user.TeamName, &user.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// FindUserByID finds a user by their ID.
func (r *UserRepositoryImpl) FindUserByID(userID string) (*db2.User, error) {
	var user db2.User
	err := r.db.QueryRow("SELECT id, username, team_name, is_active FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.TeamName, &user.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
