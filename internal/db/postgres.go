package db

import (
	"AvitoPRService/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

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
		log.Fatalf(err.Error())
	}
	return conn
}
