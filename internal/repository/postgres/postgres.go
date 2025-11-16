package postgres

import (
	"AvitoPRService/internal/config"
	"database/sql"
	"fmt"
	"log"
)

// CreateConnection creates a new connection to the database
func CreateConnection(config *config.Config) *sql.DB {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
	)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn
}
