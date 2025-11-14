package repository

import (
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/model"
)

// TeamRepository is an interface for working with teams in the database
type TeamRepository interface {
	CreateTeam(teamName string, members []dto.TeamMemberDto) (*model.Team, error)
	IsTeamExists(teamName string) (bool, error)
	FindTeamByName(teamName string) (*model.Team, error)
}
