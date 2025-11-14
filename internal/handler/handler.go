package handler

import "AvitoPRService/internal/service"

type Handler struct {
	UserHandler *UserHandler
	TeamHandler *TeamHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(service.UserService),
		TeamHandler: NewTeamHandler(service.TeamService),
	}
}
