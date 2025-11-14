// Package dto contains data transfer objects (DTOs) used for API request and response structures.
package dto

// TeamCreateDto represents the request body for creating a new team.
type TeamCreateDto struct {
	TeamName string          `json:"team_name" binding:"required"`
	Members  []TeamMemberDto `json:"members" binding:"required"`
}

// TeamMemberDto represents a member of a team.
type TeamMemberDto struct {
	ID       string `json:"user_id" binding:"required"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

// TeamGetByNameDto represents the request body for getting a team by name.
type TeamGetByNameDto struct {
	TeamName string `form:"team_name" binding:"required"`
}
