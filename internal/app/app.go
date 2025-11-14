// Package app - main app structure
package app

import (
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/handler"
	"AvitoPRService/internal/repository"
	"AvitoPRService/internal/service"
	"database/sql"
)

// App - main app structure
type App struct {
	DB *sql.DB

	Config *config.Config

	Repositories *repository.Repository
	Services     *service.Service
	Handlers     *handler.Handler
}

// NewApp - create new app instance
func NewApp(config *config.Config) *App {
	dbConn := db.CreateConnection(config)
	repositories := repository.NewRepository(dbConn)
	services := service.NewService(repositories)
	return &App{
		DB: dbConn,

		Config: config,

		Repositories: repositories,
		Services:     services,
		Handlers:     handler.NewHandler(services),
	}
}
