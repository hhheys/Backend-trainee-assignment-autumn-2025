package handler

import (
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/response"
	errorResponse "AvitoPRService/internal/response/error_response"
	"AvitoPRService/internal/security"
	"AvitoPRService/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handler for user requests
type UserHandler struct {
	s service.UserService
}

// NewUserHandler creates new UserHandler
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s: s}
}

// SetIsActive sets is_active field of user
func (h *UserHandler) SetIsActive(c *gin.Context) {
	var userSetIsActiveDto dto.UserSetIsActiveDto
	err := c.ShouldBindJSON(&userSetIsActiveDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, "validation error"))
		return
	}

	user, err := h.s.SetIsActive(userSetIsActiveDto.UserID, userSetIsActiveDto.IsActive)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.NotFound, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.NewUserResponse(user))
}

// GetAccessToken Временная версия, "заглушка". В проде реализовать адекватную авторизацию.
// GetAccessToken returns access token for user
func (h *UserHandler) GetAccessToken(c *gin.Context) {
	var userGetAccessTokenDto dto.UserGetAccessTokenDto
	err := c.ShouldBindJSON(&userGetAccessTokenDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, "validation error"))
		return
	}
	token, err := security.GenerateJWT(userGetAccessTokenDto.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.InternalServerError, err.Error()))
		return
	}
	c.SetCookie("access_token", token, 3600, "/", "", false, true)
	c.AbortWithStatus(http.StatusAccepted)
}
