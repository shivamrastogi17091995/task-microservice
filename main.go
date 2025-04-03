package main

import (
	"log"
	"os"
	"task-microservice/config"
	"task-microservice/db"
	"task-microservice/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	port := os.Getenv("PORT")
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
