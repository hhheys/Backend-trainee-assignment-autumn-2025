// Package db provides functionality for database connections,
// migrations, and error handling within the application.
package db

import "errors"

var (
	// ErrUserNotFound is returned when a user is not found.
	ErrUserNotFound = errors.New("resource not found")
	// ErrUserNoTeamFound is returned when a user has no team.
	ErrUserNoTeamFound = errors.New("resource not found")
	// ErrUserIsNotActive is returned when user have isActive = false
	ErrUserIsNotActive = errors.New("user is not active")
	// ErrTeamExists is returned when a team already exists.
	ErrTeamExists = errors.New("team %s already exists")
	// ErrTeamNotExists is returned when a team does not exist.
	ErrTeamNotExists = errors.New("resource not found")
	// ErrPullRequestAlreadyExists is returned when a pull request already exists.
	ErrPullRequestAlreadyExists = errors.New("PR id already exists")
	// ErrPullRequestNotExists is returned when a pull request does not exist.
	ErrPullRequestNotExists = errors.New("resource not found")
	// ErrPullRequestMergedReassign is returned when a pull request is already merged.
	ErrPullRequestMergedReassign = errors.New("cannot reassign on merged PR")
	// ErrUserIsNotAssignedToPR is returned when a user is not assigned to a pull request.
	ErrUserIsNotAssignedToPR = errors.New("reviewer is not assigned to this PR")
	// ErrNoActiveReplacementCandidates occurs when no active replacement candidates exist in a team.
	ErrNoActiveReplacementCandidates = errors.New("no active replacement candidate in team")
)
