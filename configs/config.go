package configs

import (
	"fmt"
	"log"
	"os"
)

var DATABASE_URL = os.Getenv("DATABASE_URL")
var SERVER_HOSTNAME = os.Getenv("SERVER_HOSTNAME")
var DOCS_HOSTNAME = os.Getenv("DOCS_HOSTNAME")
var GOOGLE_CALLBACK_DOMAIN = os.Getenv("GOOGLE_CALLBACK_DOMAIN")
var GOOGLE_OAUTH_CLIENT_ID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
var GOOGLE_OAUTH_CLIENT_SECRET = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")

func LoadAndCheck() {
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	if SERVER_HOSTNAME == "" {
		log.Fatal("SERVER_HOSTNAME not set")
	}

	if DOCS_HOSTNAME == "" {
		log.Fatal("DOCS_HOSTNAMENAME not set")
	}

	if GOOGLE_CALLBACK_DOMAIN == "" {
		log.Fatal("GOOGLE_CALLBACK_DOMAIN not set")
	}

	if GOOGLE_OAUTH_CLIENT_ID == "" {
		log.Fatal("GOOGLE_OAUTH_CLIENT_ID not set")
	}

	if GOOGLE_OAUTH_CLIENT_SECRET == "" {
		log.Fatal("GOOGLE_OAUTH_CLIENT_SECRET not set")
	}

	fmt.Println("1", "DATABASE_URL", "=", DATABASE_URL)
	fmt.Println("2", "SERVER_HOSTNAME", "=", SERVER_HOSTNAME)
	fmt.Println("3", "DOCS_HOSTNAME", "=", DOCS_HOSTNAME)
	fmt.Println("4", "GOOGLE_CALLBACK_DOMAIN", "=", GOOGLE_CALLBACK_DOMAIN)
	fmt.Println("5", "GOOGLE_OAUTH_CLIENT_ID", "=", GOOGLE_OAUTH_CLIENT_ID)
	fmt.Println("6", "GOOGLE_OAUTH_CLIENT_SECRET", "=", GOOGLE_OAUTH_CLIENT_SECRET)
}
