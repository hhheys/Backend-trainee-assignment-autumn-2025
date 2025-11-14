package response

import "AvitoPRService/internal/model"

type TeamCreateResponse struct {
	TeamName string               `json:"team_name" binding:"required"`
	Members  []TeamMemberResponse `json:"members" binding:"required"`
}

type TeamMemberResponse struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func NewTeamMemberResponse(users []model.User) []TeamMemberResponse {
	result := make([]TeamMemberResponse, len(users))
	for i, user := range users {
		result[i] = TeamMemberResponse{
			ID:       user.ID,
			Username: user.Username,
			IsActive: user.IsActive,
		}
	}
	return result
}

func NewTeamCreateResponse(team *model.Team) *TeamCreateResponse {
	return &TeamCreateResponse{
		TeamName: team.TeamName,
		Members:  NewTeamMemberResponse(team.Members),
	}
}
