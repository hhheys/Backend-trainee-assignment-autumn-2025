package user

import "AvitoPRService/internal/model"

type UserService interface {
	SetIsActive(userId uint, isActive bool) (*model.User, error)
}
