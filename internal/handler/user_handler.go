package handler

import (
	"AvitoPRService/internal/service/user"
)

type UserHandler struct {
	s user.UserService
}

func NewUserHandler(s user.UserService) *UserHandler {
	return &UserHandler{s: s}
}
