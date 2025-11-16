package handler

import (
	"AvitoPRService/internal/model/dto"
	"AvitoPRService/internal/model/response"
	response2 "AvitoPRService/internal/model/response/error_response"
	"AvitoPRService/internal/repository/postgres"
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
	if err := c.ShouldBindJSON(&teamCreateDto); err != nil {
		if status, errResp := response2.HandleError(err); errResp != nil {
			c.JSON(status, errResp)
			return
		}
	}

	createTeam, err := h.r.CreateTeam(teamCreateDto.TeamName, teamCreateDto.Members)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	c.JSON(http.StatusCreated, response.NewTeamCreateResponse(createTeam))
}

// FindByName finds a team by name
func (h *TeamHandler) FindByName(c *gin.Context) {
	var teamGetByNameDto dto.TeamGetByNameDto
	err := c.ShouldBindQuery(&teamGetByNameDto)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	foundedTeam, err := h.r.FindTeamByName(teamGetByNameDto.TeamName)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(foundedTeam))
}
