package user

import "AvitoPRService/internal/model"

type UserRepository interface {
	SetIsActive(userId string, isActive bool) (*model.User, error)
}
