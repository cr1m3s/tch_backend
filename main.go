package main

import (
	"log"
	"net/http"

	"github.com/cr1m3s/tch_backend/controllers"
	"github.com/cr1m3s/tch_backend/config"
	_ "github.com/cr1m3s/tch_backend/docs"
	middleware "github.com/cr1m3s/tch_backend/middlewares"
	"github.com/cr1m3s/tch_backend/models"
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
	
	server := gin.Default()
	// cors.Default() allows all origins
	server.Use(cors.Default())

	// router.POST("/register", controllers.Register)
	// localhost gonna be used by default
	AuthController := controllers.NewUsersController()
	AuthGoogleController := controllers.NewAuthGoogleController()

	router := server.Group("/api")
	router.GET("/", HealthCheck)
	url := ginSwagger.URL(config.GetConfig().DOCS_HOSTNAME + "/api/docs/doc.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/auth/register", AuthController.UserRegister)
	router.POST("/auth/login", AuthController.UserLogin)
	router.GET("/auth/login-google", AuthGoogleController.LoginGoogle)
	router.GET("/auth/login-google-callback", AuthGoogleController.LoginGoogleCallback)
	protected := server.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/userinfo", AuthController.UserInfo)

	log.Fatal(server.Run(config.GetConfig().SERVER_HOSTNAME))
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
	c.JSON(http.StatusOK, models.NewResponseSuccess("Server up and running."))
}
