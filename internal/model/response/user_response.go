package response

import (
	"AvitoPRService/internal/model/db"
)

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

// UserAssignedReviewsResponse represents the response for assigned reviews.
type UserAssignedReviewsResponse struct {
	UserID       string                 `json:"user_id"`
	PullRequests []*PullRequestResponse `json:"pull_requests"`
}

// NewUserResponse converts a model.User to a UserWrapperResponse.
func NewUserResponse(user *db.User) *UserWrapperResponse {
	return &UserWrapperResponse{
		User: &UserResponse{
			ID:       user.ID,
			Username: user.Username,
			TeamName: user.TeamName,
			IsActive: user.IsActive,
		},
	}
}

// NewUserAssignedReviewsResponse creates a new UserAssignedReviewsResponse.
func NewUserAssignedReviewsResponse(userID string, pullRequests []*db.PullRequest) *UserAssignedReviewsResponse {
	return &UserAssignedReviewsResponse{
		UserID:       userID,
		PullRequests: NewPullRequestResponses(pullRequests),
	}
}
