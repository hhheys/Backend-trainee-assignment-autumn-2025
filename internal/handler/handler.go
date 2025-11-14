// Package handler provides HTTP request handlers for the application's API,
// processing input, managing responses, and connecting services to routes.
package handler

import "AvitoPRService/internal/service"

// Handler provides handlers for the service.
type Handler struct {
	UserHandler *UserHandler
	TeamHandler *TeamHandler
}

// NewHandler returns a new Handler.
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(services.UserService),
		TeamHandler: NewTeamHandler(services.TeamService),
	}
}
