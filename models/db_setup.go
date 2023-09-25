package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"database/sql"
	_ "github.com/lib/pq" 
)


func ConnectDataBase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		dbname	 = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
		dbdriver = os.Getenv("DB_DRIVER")
	)
	
	dbName := "golang_postgres"
	dburl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	host,
	port,
	user,
	dbname,
	password,
	)

	DB, err := sql.Open(dbdriver, dburl)

	if err != nil {
		fmt.Println("Cannot connect to database ", dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Connected to the database ", dbdriver)
	}
	existsQuery := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s'", dbName)
	var exists bool
	err = DB.QueryRow(existsQuery).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("Database '%s' already exists\n", dbName)
	} else {
		_, err = DB.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			log.Fatal(err)
		}
	}

	var extensionExists bool
	err = DB.QueryRow("SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp'").Scan(&extensionExists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if !extensionExists {
		// Create the extension "uuid-ossp" since it doesn't exist.
		_, err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Extension 'uuid-ossp' created")
	} else {
		fmt.Println("Extension 'uuid-ossp' already exists")
	}
	return DB
}
