package config

import (
	"os"
	"sync"
	"log"
	"reflect"
)

type AppConfig struct {
	SERVER_HOSTNAME string
	DOCS_HOSTNAME string
	DATABASE_URL string
} 
  
var conf *AppConfig
var once sync.Once

func GetConfig() *AppConfig {
	once.Do(func() {		
		conf = &AppConfig{ 
			SERVER_HOSTNAME: os.Getenv("SERVER_HOSTNAME"),
			DOCS_HOSTNAME: os.Getenv("DOCS_HOSTNAME"),
			DATABASE_URL: os.Getenv("DATABASE_URL"),
		}
		
		values := reflect.ValueOf(*conf)
		types := values.Type()

		for i := 0; i < values.NumField(); i++ {
			if values.Field(i).IsZero() {
				log.Fatal("env ", types.Field(i).Name, " is empty.")
			}
		}
	})
	 
	return conf 
} 