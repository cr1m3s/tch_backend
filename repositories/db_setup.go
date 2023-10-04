package repositories

import (
	"context"
	"log"

	"github.com/cr1m3s/tch_backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDataBase() *pgxpool.Pool {

	dbpool, err := pgxpool.New(context.Background(), config.GetConfig().DATABASE_URL)
	if err != nil {
		log.Fatal("create database conection error:" + err.Error())
	}

	return dbpool
}
