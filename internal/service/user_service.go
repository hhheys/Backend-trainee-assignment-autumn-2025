package service

import (
	"AvitoPRService/internal/model"
	"AvitoPRService/internal/repository"
)

// UserService defines the business operations related to teams.
type UserService interface {
	SetIsActive(userID string, isActive bool) (*model.User, error)
}

// UserServiceImpl is a implementation of Service
type UserServiceImpl struct {
	Repository repository.UserRepository
}

// NewUserServiceImpl creates a copy of UserServiceImpl
func NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{Repository: repository}
}

// SetIsActive changes user and sets is_active attribute
func (s *UserServiceImpl) SetIsActive(userID string, isActive bool) (*model.User, error) {
	user, err := s.Repository.SetIsActive(userID, isActive)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
