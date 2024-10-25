package postgres

import (
	"fullstack-snippetbox-backend/internal/config"
	"log"

	"github.com/jackc/pgx"
)

var database *pgx.Conn

func init() {
	appConfig := config.GetAppConfig()

	conn, err := pgx.Connect(pgx.ConnConfig{Host: "hui"})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	database = conn
}

func GetDB() *pgx.Conn  {
	return database
}
