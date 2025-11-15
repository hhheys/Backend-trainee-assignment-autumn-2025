// PR Service for Avito backend trainee assignment 2025
package main

import (
	"AvitoPRService/internal/app"
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/config/env"
	"AvitoPRService/internal/config/logger"
	"AvitoPRService/internal/router"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	logger.LogInit()
	logger.Logger.Info("Starting Avito-PR backend server")

	env.LoadEnv()
	appConfig := config.NewConfig()
	application := app.NewApp(appConfig)

	r := router.NewRouter(application)

	err := r.Run(fmt.Sprintf("localhost:%d", appConfig.ServerPort))
	if err != nil {
		logger.Logger.Fatalf("Failed to start an application with error %s", err.Error())
	}
	logger.Logger.Info("Application started successfully")
}
