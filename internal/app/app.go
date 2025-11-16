// Package app - main app structure
package app

import (
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/handler"
	"AvitoPRService/internal/repository/postgres"
	"database/sql"
)

// App - main app structure
type App struct {
	DB *sql.DB

	Config *config.Config

	Repositories *postgres.Repository
	Handlers     *handler.Handler
}

// NewApp - create new app instance
func NewApp(config *config.Config) *App {
	dbConn := postgres.CreateConnection(config)
	repositories := postgres.NewRepository(dbConn)
	return &App{
		DB: dbConn,

		Config: config,

		Repositories: repositories,
		Handlers:     handler.NewHandler(repositories),
	}
}
