// Package env provides structures and utilities for reading
// and managing application environment variables and configuration.
package env

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv Loads env variables from .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env files found")
	}
}
