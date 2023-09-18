package main

import (
	"net/http"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/cr1m3s/tch_backend/docs/ginsimple"
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
	db_url := os.Getenv("DATABASE_URL")
	fmt.Println(db_url)

	router := gin.New()
	
	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
        
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/", HealthCheck)
	
	router.Run("localhost:8000")
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
