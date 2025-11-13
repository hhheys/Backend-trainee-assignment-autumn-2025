package handler

import (
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/response"
	errorResponse "AvitoPRService/internal/response/error_response"
	"AvitoPRService/internal/service/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	s user.UserService
}

func NewUserHandler(s user.UserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) SetIsActive(c *gin.Context) {
	var userSetIsActiveDto dto.UserSetIsActiveDto
	err := c.ShouldBindJSON(&userSetIsActiveDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BAD_REQUEST, "validation error"))
		return
	}

	user, err := h.s.SetIsActive(userSetIsActiveDto.UserId, userSetIsActiveDto.IsActive)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.NOT_FOUND, err.Error()))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.INTERNAL_SERVER_ERROR, err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, response.NewUserResponse(user))
}
