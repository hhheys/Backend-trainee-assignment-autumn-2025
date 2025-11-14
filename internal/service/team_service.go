package service

import (
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/model"
	"AvitoPRService/internal/repository"
)

// TeamService defines the business operations related to teams.
type TeamService interface {
	CreateTeam(teamCreateDto *dto.TeamCreateDto) (*model.Team, error)
	FindTeamByName(teamName string) (*model.Team, error)
}

// TeamServiceImpl is a implementation of TeamService
type TeamServiceImpl struct {
	Repository repository.TeamRepository
}

// NewTeamServiceImpl creates new copy of TeamServiceImpl
func NewTeamServiceImpl(repository repository.TeamRepository) *TeamServiceImpl {
	return &TeamServiceImpl{Repository: repository}
}

// CreateTeam creates new team and returns model.Team
func (s *TeamServiceImpl) CreateTeam(teamCreateDto *dto.TeamCreateDto) (*model.Team, error) {
	createdTeam, err := s.Repository.CreateTeam(teamCreateDto.TeamName, teamCreateDto.Members)
	if err != nil {
		return nil, err
	}
	return createdTeam, nil
}

// FindTeamByName finds team by name if it's exists. Else throw an error
func (s *TeamServiceImpl) FindTeamByName(teamName string) (*model.Team, error) {
	foundTeam, err := s.Repository.FindTeamByName(teamName)
	if err != nil {
		return nil, err
	}
	return foundTeam, nil
}
