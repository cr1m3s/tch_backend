package main

import (
	"log"

	"github.com/cr1m3s/tch_backend/configs"
	"github.com/cr1m3s/tch_backend/routers"
)

// @title						Study marketplace API
// @version						0.0.1
// @description					Marketplace to connect students and teachers
// @termsOfService				[TODO]
// @contact.name				API Support
// @contact.url					[TODO]
// @contact.email				[TODO]
// @license.name				[TODO]
// @license.url					[TODO]
// @BasePath					/
// @schemes						http https
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
func main() {
	configs.LoadAndCheck()
	server := routers.SetupRouter()
	routers.SetupCORS(server)

	log.Fatal(server.Run(configs.SERVER_HOSTNAME))
}

