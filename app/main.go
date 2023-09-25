package main

import (
	"log"
	"net/http"

	"github.com/cr1m3s/tch_backend/app/controllers"
	"github.com/cr1m3s/tch_backend/app/db"
	dbConn "github.com/cr1m3s/tch_backend/app/db/sqlc"
	"github.com/cr1m3s/tch_backend/app/middleware"
	_ "github.com/cr1m3s/tch_backend/app/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbConn.Queries

	AuthController controllers.AuthController
)

//	@title						Study marketplace API
//	@version					0.0.1
//	@description				Marketplace to connect students and teachers
//	@termsOfService				[TODO]
//	@contact.name				API Support
//	@contact.url				[TODO]
//	@contact.email				[TODO]
//	@license.name				[TODO]
//	@license.url				[TODO]
//	@host						localhost:8000
//	@BasePath					/
//	@schemes					http
//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
func main() {
	// DATABASE_URL=postgresql://USER:PASSWORD@HOST:PORT/DATABASE
	conn := models.ConnectDataBase()
	db = dbConn.New(conn)
	server := gin.Default()
	// cors.Default() allows all origins
	server.Use(cors.Default())

	// router.POST("/register", controllers.Register)
	// localhost gonna be used by default
	AuthController = *controllers.NewAuthController(db)

	router := server.Group("/api")

	router.GET("/", HealthCheck)
	url := ginSwagger.URL("http://localhost:8000/api/docs/doc.json")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/auth/register", AuthController.SignUpUser)
	router.POST("/auth/login", AuthController.LoginUser)

	protected := server.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/userinfo", AuthController.GetUserInfo)

	log.Fatal(server.Run(":8000"))
}

//  HealthCheck godoc
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/api/ [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and runing",
	}

	c.JSON(http.StatusOK, res)
}
