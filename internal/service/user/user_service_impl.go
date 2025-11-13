package user

import (
	"AvitoPRService/internal/model"
	"AvitoPRService/internal/repository/user"
)

type UserServiceImpl struct {
	Repository user.UserRepository
}

func NewUserServiceImpl(repository user.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{Repository: repository}
}

func (s *UserServiceImpl) SetIsActive(userId string, isActive bool) (*model.User, error) {
	user, err := s.Repository.SetIsActive(userId, isActive)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
