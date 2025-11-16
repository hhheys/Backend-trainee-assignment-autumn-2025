package dto

// UserSetIsActiveDto is a DTO for setting the is_active field of a user.
type UserSetIsActiveDto struct {
	UserID   string `json:"user_id" binding:"required"`
	IsActive bool   `json:"is_active"`
}

// UserGetAccessTokenDto is a DTO for getting the access token of a user.
type UserGetAccessTokenDto struct {
	UserID string `json:"user_id" binding:"required"`
}

// UserGetPRsDto is a DTO for getting the PRs of a user.
type UserGetPRsDto struct {
	UserID string `form:"user_id" binding:"required"`
}
