package main

import (
	"log"
	"net/http"

	"github.com/cr1m3s/tch_backend/configs"
	"github.com/cr1m3s/tch_backend/controllers"
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
	configs.LoadAndCheck()
	server := gin.Default()
	sorsConfig := cors.DefaultConfig()
	sorsConfig.AddAllowHeaders("Access-Control-Allow-Headers")
	sorsConfig.AddAllowHeaders("Access-Control-Request-Method")
	sorsConfig.AddAllowHeaders("Access-Control-Request-Headers")
	sorsConfig.AddAllowHeaders("X-Requested-With")
	sorsConfig.AddAllowHeaders("Accept")
	sorsConfig.AddAllowHeaders("Authorization")
	sorsConfig.AllowAllOrigins = true
	sorsConfig.AllowCredentials = true
	c := cors.New(sorsConfig)
	server.Use(c)

	// router.POST("/register", controllers.Register)
	// localhost gonna be used by default
	AuthController := controllers.NewUsersController()
	AuthGoogleController := controllers.NewAuthGoogleController()

	router := server.Group("/api")
	router.GET("/", HealthCheck)
	url := ginSwagger.URL(configs.DOCS_HOSTNAME + "/api/docs/doc.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/auth/register", AuthController.UserRegister)
	router.POST("/auth/login", AuthController.UserLogin)
	router.GET("/auth/login-google", AuthGoogleController.LoginGoogle)
	router.GET("/auth/login-google-callback", AuthGoogleController.LoginGoogleCallback)
	protected := server.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/userinfo", AuthController.UserInfo)

	log.Fatal(server.Run(configs.SERVER_HOSTNAME))
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
