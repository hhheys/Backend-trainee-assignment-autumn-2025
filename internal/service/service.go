// Package service contains business logic and application services,
// implementing key operations and interactions between repositories and handlers.
package service

import (
	"AvitoPRService/internal/repository"
)

// Service defines the business operations related to application.
type Service struct {
	UserService UserService
	TeamService TeamService
}

// NewService creates a copy of Service
func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserServiceImpl(repository.UserRepository),
		TeamService: NewTeamServiceImpl(repository.TeamRepository),
	}
}
