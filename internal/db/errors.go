// Package db provides functionality for database connections,
// migrations, and error handling within the application.
package db

import "errors"

var (
	// ErrUserNotFound is returned when a user is not found.
	ErrUserNotFound = errors.New("resource not found")
	// ErrTeamExists is returned when a team already exists.
	ErrTeamExists = errors.New("team %s already exists")
	// ErrTeamNotExists is returned when a team does not exist.
	ErrTeamNotExists = errors.New("resource not found")
)
