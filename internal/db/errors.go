package db

import "errors"

var (
	ErrUserNotFound  = errors.New("resource not found")
	ErrTeamExists    = errors.New("team %s already exists")
	ErrTeamNotExists = errors.New("resource not found")
)
