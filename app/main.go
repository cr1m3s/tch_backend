package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	
	db_url := os.Getenv("DATABASE_URL")
	fmt.Println(db_url)

	router := gin.Default()
	router.Run("localhost:8000")
}


