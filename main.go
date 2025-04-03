package main

import (
	"log"
	"os"
	"task-microservice/config"
	"task-microservice/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	router := gin.Default()
	port := os.Getenv("PORT")
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
