package postgres

import (
	"AvitoPRService/internal/model/db"
	"AvitoPRService/internal/model/dto"
)

// TeamRepository is an interface for working with teams in the database
type TeamRepository interface {
	CreateTeam(teamName string, members []dto.TeamMemberDto) (*db.Team, error)
	IsTeamExists(teamName string) (bool, error)
	FindTeamByName(teamName string) (*db.Team, error)
}
