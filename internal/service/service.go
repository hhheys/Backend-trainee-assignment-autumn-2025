package service

import (
	"AvitoPRService/internal/repository"
	teamService "AvitoPRService/internal/service/team"
	userService "AvitoPRService/internal/service/user"
)

type Service struct {
	UserService userService.UserService
	TeamService teamService.TeamService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: userService.NewUserServiceImpl(repository.UserRepository),
		TeamService: teamService.NewTeamServiceImpl(repository.TeamRepository),
	}
}
