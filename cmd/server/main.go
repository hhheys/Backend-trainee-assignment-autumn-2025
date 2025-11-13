package main

import (
	"AvitoPRService/internal/app"
	"AvitoPRService/internal/config"
	"AvitoPRService/internal/router"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env files found")
	}
	config := config.NewConfig()
	app := app.NewApp(config)

	r := router.NewRouter(app)

	r.Run(fmt.Sprintf("localhost:%d", config.ServerPort))
}
