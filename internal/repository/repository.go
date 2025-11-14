package repository

import (
	"AvitoPRService/internal/repository/team"
	"AvitoPRService/internal/repository/user"
	"database/sql"
)

type Repository struct {
	UserRepository user.UserRepository
	TeamRepository team.TeamRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepositoryImpl(db),
		TeamRepository: team.NewTeamRepositoryImpl(db),
	}
}
