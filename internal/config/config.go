// Package config provides application configuration structures and
// utilities for loading, validating, and managing configuration settings.
package config

import (
	"log"
	"os"
	"strconv"
)

// Config is a struct for configuration
type Config struct {
	// Data for database
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string

	// Service port
	ServerPort int

	// Secret string for JWT
	SecretString string

	// Access token for admin permissions
	AccessToken string
}

// NewConfig creates a new Config
func NewConfig() *Config {
	portDBStr := getEnvOrFatal("DATABASE_PORT")
	portDB, err := strconv.Atoi(portDBStr)
	if err != nil {
		log.Fatalf("Invalid DATABASE_PORT: %v", err)
	}

	portServerStr := getEnvOrFatal("SERVER_PORT")
	portServer, err := strconv.Atoi(portServerStr)
	if err != nil {
		log.Fatalf("Invalid SERVER_PORT: %v", err)
	}

	return &Config{
		DatabasePort:     portDB,
		DatabaseUser:     getEnvOrFatal("DATABASE_USER"),
		DatabasePassword: getEnvOrFatal("DATABASE_PASSWORD"),
		DatabaseName:     getEnvOrFatal("DATABASE_NAME"),
		DatabaseHost:     getEnvOrFatal("DATABASE_HOST"),
		ServerPort:       portServer,
		SecretString:     getEnvOrFatal("SECRET_STRING"),
		AccessToken:      getEnvOrFatal("ACCESS_TOKEN"),
	}
}

func getEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Empty environment variable: %s", key)
	}
	return value
}
