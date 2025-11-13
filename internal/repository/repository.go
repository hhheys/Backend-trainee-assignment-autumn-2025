package repository

import (
	"AvitoPRService/internal/repository/user"
	"database/sql"
)

type Repository struct {
	UserRepository user.UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: user.NewUserRepositoryImpl(db),
	}
}
