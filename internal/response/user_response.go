package response

import "AvitoPRService/internal/model"

type UserWrapperResponse struct {
	User *UserResponse `json:"user"`
}

type UserResponse struct {
	ID       uint    `json:"user_id"`
	Username string  `json:"username"`
	TeamName *string `json:"team_name"`
	IsActive bool    `json:"is_active"`
}

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
