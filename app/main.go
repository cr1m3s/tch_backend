package main

import (
	"net/http"

	_ "github.com/cr1m3s/tch_backend/docs/ginsimple"
//	"github.com/cr1m3s/tch_backend/app/controllers"
	"github.com/cr1m3s/tch_backend/app/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Study marketplace API
// @version 0.0.1
// @description Marketplace to connect students and teachers
// @termsOfService [TODO]

// @contact.name API Support
// @contact.url [TODO]
// @contact.email [TODO]

// @license.name [TODO]
// @license.url [TODO]

// @host localhost:8000
// @BasePath /
// @schemes http

func main() {

	// DATABASE_URL=postgresql://USER:PASSWORD@HOST:PORT/DATABASE
	models.ConnectDataBase()

	router := gin.Default()
	// cors.Default() allows all origins
	router.Use(cors.Default())

	url := ginSwagger.URL("http://localhost:8000/docs/doc.json")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/", HealthCheck)
	
	// router.POST("/register", controllers.Register)
	// localhost gonna be used by default
	router.Run(":8000")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and runing",
	}

	c.JSON(http.StatusOK, res)
}
