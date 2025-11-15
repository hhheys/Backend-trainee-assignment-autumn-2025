// Package repository contains interfaces and implementations
// for data access layers, enabling abstraction over database operations for application entities.
package repository

import (
	"database/sql"
)

// Repository provides access to the repository.
type Repository struct {
	UserRepository        UserRepository
	TeamRepository        TeamRepository
	PullRequestRepository PullRequestRepository
}

// NewRepository creates a new Repository.
func NewRepository(db *sql.DB) *Repository {
	userRepository := NewUserRepositoryImpl(db)
	teamRepository := NewTeamRepositoryImpl(db)

	return &Repository{
		UserRepository:        userRepository,
		TeamRepository:        teamRepository,
		PullRequestRepository: NewPullRequestRepositoryImpl(db, userRepository, teamRepository),
	}
}
