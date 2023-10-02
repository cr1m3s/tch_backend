package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cr1m3s/tch_backend/controllers"
	_ "github.com/cr1m3s/tch_backend/docs"
	middleware "github.com/cr1m3s/tch_backend/middlewares"
	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/repositories"
	"github.com/cr1m3s/tch_backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("env DATABASE_URL is empty")
	}

	server_host := os.Getenv("SERVER_HOSTNAME")
	if server_host == "" {
		log.Fatal("env SERVER_HOSTNAME is empty")
	}

	docs_host := os.Getenv("DOCS_HOSTNAME")
	if docs_host == "" {
		log.Fatal("env DOCS_HOSTNAME is empty")
	}

	db := repositories.NewAppRepository()
	server := gin.Default()
	// cors.Default() allows all origins
	server.Use(cors.Default())

	// router.POST("/register", controllers.Register)
	// localhost gonna be used by default
	AuthController := controllers.NewUsersController(db)
	UserService := services.NewUserService(db)
	AuthGoogleController := controllers.NewAuthGoogleController(UserService)

	router := server.Group("/api")
	router.GET("/", HealthCheck)
	url := ginSwagger.URL(docs_host + "/api/docs/doc.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/auth/register", AuthController.UserRegister)
	router.POST("/auth/login", AuthController.UserLogin)
	router.GET("/auth/login-google", AuthGoogleController.LoginGoogle)
	router.GET("/auth/login-google-callback", AuthGoogleController.LoginGoogleCallback)
	protected := server.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/userinfo", AuthController.UserInfo)

	log.Fatal(server.Run(server_host))
}

// HealthCheck godoc
// @Summary		Show the status of server.
// @Description	get the status of server.
// @Tags			root
// @Accept			*/*
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/api/ [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{Status: "succes", Data: "Server up and running."})
}
