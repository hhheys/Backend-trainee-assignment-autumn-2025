// Package handler provides HTTP request handlers for the application'r API,
// processing input, managing responses, and connecting services to routes.
package handler

import (
	"AvitoPRService/internal/repository/postgres"
)

// Handler provides handlers for the service.
type Handler struct {
	UserHandler        *UserHandler
	TeamHandler        *TeamHandler
	PullRequestHandler *PullRequestHandler
}

// NewHandler returns a new Handler.
func NewHandler(repositories *postgres.Repository) *Handler {
	return &Handler{
		UserHandler:        NewUserHandler(repositories.UserRepository),
		TeamHandler:        NewTeamHandler(repositories.TeamRepository),
		PullRequestHandler: NewPullRequestHandler(repositories.PullRequestRepository),
	}
}
