package team

import (
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/model"
)

type TeamService interface {
	CreateTeam(teamCreateDto *dto.TeamCreateDto) (*model.Team, error)
	FindTeamByName(teamName string) (*model.Team, error)
}
