package dto

type UserSetIsActiveDto struct {
	UserId   uint `json:"user_id" binding:"required"`
	IsActive bool `json:"is_active"`
}

type UserGetAccessTokenDto struct {
	UserId uint `json:"user_id" binding:"required"`
}
