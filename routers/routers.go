package routers

import (
	"github.com/cr1m3s/tch_backend/configs"
	"github.com/cr1m3s/tch_backend/controllers"
	_ "github.com/cr1m3s/tch_backend/docs"
	middleware "github.com/cr1m3s/tch_backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(server *gin.Engine) *gin.Engine {

	AuthController := controllers.NewUsersController()
	AuthGoogleController := controllers.NewAuthGoogleController()
	docs_url := ginSwagger.URL(configs.DOCS_HOSTNAME + "/api/docs/doc.json")

	api := server.Group("/api")

	api.GET("/", controllers.HealthCheck)

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docs_url))
	api.POST("/auth/register", AuthController.UserRegister)
	api.POST("/auth/login", AuthController.UserLogin)
	api.GET("/auth/login-google", AuthGoogleController.LoginGoogle)
	api.GET("/auth/login-google-callback", AuthGoogleController.LoginGoogleCallback)

	protected := server.Group("/protected")

	protected.Use(middleware.AuthMiddleware())
	protected.GET("/userinfo", AuthController.UserInfo)
	protected.PATCH("/user-patch", AuthController.UserPatch)

	return server
}

func SetupCORS(server *gin.Engine) *gin.Engine {

	sorsConfig := cors.DefaultConfig()
	sorsConfig.AllowAllOrigins = true
	sorsConfig.AllowCredentials = true
	sorsConfig.AddAllowHeaders("Access-Control-Allow-Headers")
	sorsConfig.AddAllowHeaders("Access-Control-Request-Method")
	sorsConfig.AddAllowHeaders("Access-Control-Request-Headers")
	sorsConfig.AddAllowHeaders("Access-Control-Allow-Origin")
	sorsConfig.AddAllowHeaders("X-Requested-With")
	sorsConfig.AddAllowHeaders("Accept")
	sorsConfig.AddAllowHeaders("Authorization")
	c := cors.New(sorsConfig)
	server.Use(c)

	return server
}
