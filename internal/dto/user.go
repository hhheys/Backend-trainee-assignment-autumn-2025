package dto

type UserSetIsActiveDto struct {
	UserId   string `json:"user_id" binding:"required"`
	IsActive bool   `json:"is_active"`
}

type UserGetAccessTokenDto struct {
	UserId string `json:"user_id" binding:"required"`
}
