package repositories

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDataBase() *sql.DB {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("env DATABASE_URL is empty")
	}

	dbpool, err := sql.Open("pgx/v5", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("unable to create connection pool:" + err.Error())
	}

	return dbpool
}
