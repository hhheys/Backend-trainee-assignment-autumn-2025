package user

import "AvitoPRService/internal/model"

type UserRepository interface {
	SetIsActive(userId uint, isActive bool) (*model.User, error)
}
