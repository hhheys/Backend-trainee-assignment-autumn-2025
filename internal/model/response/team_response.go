// Package response provides types for API responses,
// including error response structures and other payloads returned by the server.
package response

import (
	"AvitoPRService/internal/model/db"
)

// TeamCreateResponse represents response for team creation
type TeamCreateResponse struct {
	TeamName string               `json:"team_name" binding:"required"`
	Members  []TeamMemberResponse `json:"members" binding:"required"`
}

// TeamMemberResponse represents response for team members
type TeamMemberResponse struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

// NewTeamMemberResponse creates new TeamMemberResponse
func NewTeamMemberResponse(users []db.User) []TeamMemberResponse {
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

// NewTeamCreateResponse creates new TeamCreateResponse
func NewTeamCreateResponse(team *db.Team) *TeamCreateResponse {
	return &TeamCreateResponse{
		TeamName: team.Name,
		Members:  NewTeamMemberResponse(team.Members),
	}
}
