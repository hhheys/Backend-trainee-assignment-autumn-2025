// PR Service for Avito backend trainee assignment 2025
package main

import (
	"AvitoPRService/internal/app"
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/router"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env files found")
	}
	config := config.NewConfig()
	app := app.NewApp(config)

	r := router.NewRouter(app)

	err := r.Run(fmt.Sprintf("localhost:%d", config.ServerPort))
	if err != nil {
		log.Fatalf("Couldn't start an application. %s", err.Error())
	}
}
