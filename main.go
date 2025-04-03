package main

import (
	"log"
	"os"
	"task-microservice/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	router := gin.Default()
	port := os.Getenv("PORT")
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
