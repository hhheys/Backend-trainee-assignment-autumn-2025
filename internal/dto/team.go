package dto

type TeamCreateDto struct {
	TeamName string          `json:"team_name" binding:"required"`
	Members  []TeamMemberDto `json:"members" binding:"required"`
}

type TeamMemberDto struct {
	ID       string `json:"user_id" binding:"required"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type TeamGetByNameDto struct {
	TeamName string `form:"team_name" binding:"required"`
}
