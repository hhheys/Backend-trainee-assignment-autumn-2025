package team

import (
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/model"
	"AvitoPRService/internal/repository/team"
)

type TeamServiceImpl struct {
	Repository team.TeamRepository
}

func NewTeamServiceImpl(repository team.TeamRepository) *TeamServiceImpl {
	return &TeamServiceImpl{Repository: repository}
}

func (s *TeamServiceImpl) CreateTeam(teamCreateDto *dto.TeamCreateDto) (*model.Team, error) {
	createdTeam, err := s.Repository.CreateTeam(teamCreateDto.TeamName, teamCreateDto.Members)
	if err != nil {
		return nil, err
	}
	return createdTeam, nil
}

func (s *TeamServiceImpl) FindTeamByName(teamName string) (*model.Team, error) {
	foundTeam, err := s.Repository.FindTeamByName(teamName)
	if err != nil {
		return nil, err
	}
	return foundTeam, nil
}
