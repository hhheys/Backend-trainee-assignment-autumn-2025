package user

import "AvitoPRService/internal/model"

type UserService interface {
	SetIsActive(userId string, isActive bool) (*model.User, error)
}
