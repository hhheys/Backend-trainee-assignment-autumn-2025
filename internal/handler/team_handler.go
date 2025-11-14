package handler

import (
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/response"
	errorResponse "AvitoPRService/internal/response/error_response"
	"AvitoPRService/internal/service/team"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	s team.TeamService
}

func NewTeamHandler(s team.TeamService) *TeamHandler {
	return &TeamHandler{s: s}
}

func (h *TeamHandler) AddTeam(c *gin.Context) {
	var teamCreateDto dto.TeamCreateDto
	err := c.ShouldBindJSON(&teamCreateDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BAD_REQUEST, "validation error"))
		return
	}
	createTeam, err := h.s.CreateTeam(&teamCreateDto)
	if err != nil {
		if errors.Is(err, db.ErrTeamExists) {
			c.JSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.NOT_FOUND, fmt.Sprintf(err.Error(), teamCreateDto.TeamName)))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.INTERNAL_SERVER_ERROR, err.Error()))
			return
		}
	}
	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(createTeam))
}

func (h *TeamHandler) FindByName(c *gin.Context) {
	var teamGetByNameDto dto.TeamGetByNameDto
	err := c.ShouldBindQuery(&teamGetByNameDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BAD_REQUEST, "validation error"))
		return
	}

	foundedTeam, err := h.s.FindTeamByName(teamGetByNameDto.TeamName)
	if err != nil {
		if errors.Is(err, db.ErrTeamNotExists) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.NOT_FOUND, err.Error()))
			return
		}
	}
	c.JSON(http.StatusAccepted, response.NewTeamCreateResponse(foundedTeam))
}
