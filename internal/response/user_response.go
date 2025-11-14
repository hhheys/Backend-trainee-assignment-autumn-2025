package response

import "AvitoPRService/internal/model"

// UserWrapperResponse represents a wrapper for the user response.
type UserWrapperResponse struct {
	User *UserResponse `json:"user"`
}

// UserResponse represents the user details in the response.
type UserResponse struct {
	ID       string  `json:"user_id"`
	Username string  `json:"username"`
	TeamName *string `json:"team_name"`
	IsActive bool    `json:"is_active"`
}

// NewUserResponse converts a model.User to a UserWrapperResponse.
func NewUserResponse(user *model.User) *UserWrapperResponse {
	return &UserWrapperResponse{
		User: &UserResponse{
			ID:       user.ID,
			Username: user.Username,
			TeamName: user.TeamName,
			IsActive: user.IsActive,
		},
	}
}
