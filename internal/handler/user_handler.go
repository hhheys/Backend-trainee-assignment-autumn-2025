package handler

import (
	"AvitoPRService/internal/model/dto"
	"AvitoPRService/internal/model/response"
	response2 "AvitoPRService/internal/model/response/error_response"
	"AvitoPRService/internal/repository/postgres"
	"AvitoPRService/internal/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handler for user requests
type UserHandler struct {
	r postgres.UserRepository
}

// NewUserHandler creates new UserHandler
func NewUserHandler(s postgres.UserRepository) *UserHandler {
	return &UserHandler{r: s}
}

// SetIsActive sets is_active field of user
func (h *UserHandler) SetIsActive(c *gin.Context) {
	var userSetIsActiveDto dto.UserSetIsActiveDto
	err := c.ShouldBindJSON(&userSetIsActiveDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, "validation error"))
		return
	}

	user, err := h.r.SetIsActive(userSetIsActiveDto.UserID, userSetIsActiveDto.IsActive)
	if err != nil {
		if status, errResp := response2.HandleError(err); errResp != nil {
			c.JSON(status, errResp)
			return
		}
	}
	c.JSON(http.StatusOK, response.NewUserResponse(user))
}

// GetReview returns user reviews
func (h *UserHandler) GetReview(c *gin.Context) {
	var userGetPRsDto dto.UserGetPRsDto
	err := c.ShouldBindQuery(&userGetPRsDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, "validation error"))
		return
	}
	reviews, err := h.r.GetUserReviews(userGetPRsDto.UserID)
	if err != nil {
		if status, errResp := response2.HandleError(err); errResp != nil {
			c.JSON(status, errResp)
			return
		}
	}

	c.JSON(http.StatusOK, response.NewUserAssignedReviewsResponse(userGetPRsDto.UserID, reviews))
	return
}

// GetAccessToken Временная версия, "заглушка". В проде реализовать адекватную авторизацию.
// GetAccessToken returns access token for user
func (h *UserHandler) GetAccessToken(c *gin.Context) {
	var userGetAccessTokenDto dto.UserGetAccessTokenDto
	err := c.ShouldBindJSON(&userGetAccessTokenDto)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	token, err := security.GenerateJWT(userGetAccessTokenDto.UserID)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	c.SetCookie("access_token", token, 3600, "/", "", false, true)
	c.AbortWithStatus(http.StatusAccepted)
}
