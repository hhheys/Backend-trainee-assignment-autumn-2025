package service

import (
	"AvitoPRService/internal/repository"
	service "AvitoPRService/internal/service/user"
)

type Service struct {
	UserService service.UserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: service.NewUserServiceImpl(repository.UserRepository),
	}
}
