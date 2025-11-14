package handler

import (
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/response"
	errorResponse "AvitoPRService/internal/response/error_response"
	"AvitoPRService/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeamHandler provides methods for working with teams
type TeamHandler struct {
	s service.TeamService
}

// NewTeamHandler creates a new TeamHandler
func NewTeamHandler(s service.TeamService) *TeamHandler {
	return &TeamHandler{s: s}
}

// AddTeam adds a new team
func (h *TeamHandler) AddTeam(c *gin.Context) {
	var teamCreateDto dto.TeamCreateDto
	err := c.ShouldBindJSON(&teamCreateDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, "validation error"))
		return
	}
	createTeam, err := h.s.CreateTeam(&teamCreateDto)
	if err != nil {
		if errors.Is(err, db.ErrTeamExists) {
			c.JSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.NotFound, fmt.Sprintf(err.Error(), teamCreateDto.TeamName)))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(createTeam))
}

// FindByName finds a team by name
func (h *TeamHandler) FindByName(c *gin.Context) {
	var teamGetByNameDto dto.TeamGetByNameDto
	err := c.ShouldBindQuery(&teamGetByNameDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, "validation error"))
		return
	}

	foundedTeam, err := h.s.FindTeamByName(teamGetByNameDto.TeamName)
	if err != nil {
		if errors.Is(err, db.ErrTeamNotExists) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.NotFound, err.Error()))
			return
		}
	}
	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(foundedTeam))
}
