package model

type User struct {
	ID       uint   `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

func NewUser(userId uint, username string, teamName string, isActive bool) *User {
	return &User{
		ID:       userId,
		Username: username,
		TeamName: teamName,
		IsActive: isActive,
	}
}
