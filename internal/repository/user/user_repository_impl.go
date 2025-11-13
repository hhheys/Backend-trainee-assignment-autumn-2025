package user

import (
	"AvitoPRService/internal/model"
	"database/sql"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) SetIsActive(userId uint, isActive bool) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("UPDATE users SET is_active = $1 WHERE id = $2 RETURNING id, username, team_name, is_active", isActive, userId).Scan(&user.ID, &user.Username, &user.TeamName, &user.IsActive)
	if err != nil {
		return &model.User{}, err
	}
	return &user, nil
}
