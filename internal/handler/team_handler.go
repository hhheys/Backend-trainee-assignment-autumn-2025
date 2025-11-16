package handler

import (
	"AvitoPRService/internal/model/dto"
	"AvitoPRService/internal/model/response"
	response2 "AvitoPRService/internal/model/response/error_response"
	"AvitoPRService/internal/repository/postgres"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeamHandler provides methods for working with teams
type TeamHandler struct {
	r postgres.TeamRepository
}

// NewTeamHandler creates a new TeamHandler
func NewTeamHandler(s postgres.TeamRepository) *TeamHandler {
	return &TeamHandler{r: s}
}

// AddTeam adds a new team
func (h *TeamHandler) AddTeam(c *gin.Context) {
	var teamCreateDto dto.TeamCreateDto
	err := c.ShouldBindJSON(&teamCreateDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, "validation error"))
		return
	}
	createTeam, err := h.r.CreateTeam(teamCreateDto.TeamName, teamCreateDto.Members)
	if err != nil {
		if errors.Is(err, postgres.ErrTeamExists) {
			c.JSON(http.StatusBadRequest, response2.NewErrorResponse(response2.NotFound, fmt.Sprintf(err.Error(), teamCreateDto.TeamName)))
			return
		}
		c.JSON(http.StatusInternalServerError, response2.NewErrorResponse(response2.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(createTeam))
}

// FindByName finds a team by name
func (h *TeamHandler) FindByName(c *gin.Context) {
	var teamGetByNameDto dto.TeamGetByNameDto
	err := c.ShouldBindQuery(&teamGetByNameDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, "validation error"))
		return
	}

	foundedTeam, err := h.r.FindTeamByName(teamGetByNameDto.TeamName)
	if err != nil {
		if errors.Is(err, postgres.ErrTeamNotExists) {
			c.JSON(http.StatusNotFound, response2.NewErrorResponse(response2.NotFound, err.Error()))
			return
		}
	}
	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(foundedTeam))
}
