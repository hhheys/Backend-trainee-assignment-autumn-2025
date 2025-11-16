package db

// User represents a user in the service
type User struct {
	ID       string  `json:"user_id"`
	Username string  `json:"username"`
	TeamName *string `json:"team_name"`
	IsActive bool    `json:"is_active"`
}

// NewUser creates a new User instance
func NewUser(userID string, username string, teamName *string, isActive bool) *User {
	return &User{
		ID:       userID,
		Username: username,
		TeamName: teamName,
		IsActive: isActive,
	}
}
