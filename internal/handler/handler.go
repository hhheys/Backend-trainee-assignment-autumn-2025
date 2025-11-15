// Package handler provides HTTP request handlers for the application's API,
// processing input, managing responses, and connecting services to routes.
package handler

import (
	"AvitoPRService/internal/repository"
	"AvitoPRService/internal/service"
)

// Handler provides handlers for the service.
type Handler struct {
	UserHandler        *UserHandler
	TeamHandler        *TeamHandler
	PullRequestHandler *PullRequestHandler
}

// NewHandler returns a new Handler.
func NewHandler(repositories *repository.Repository, services *service.Service) *Handler {
	return &Handler{
		UserHandler:        NewUserHandler(services.UserService),
		TeamHandler:        NewTeamHandler(services.TeamService),
		PullRequestHandler: NewPullRequestHandler(repositories.PullRequestRepository),
	}
}
