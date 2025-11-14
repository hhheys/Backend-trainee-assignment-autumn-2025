package team

import (
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/model"
)

type TeamRepository interface {
	CreateTeam(teamName string, members []dto.TeamMemberDto) (*model.Team, error)
	IsTeamExists(teamName string) (bool, error)
	FindTeamByName(teamName string) (*model.Team, error)
}
