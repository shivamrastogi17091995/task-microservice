package main

import (
	"log"
	"os"
	"task-microservice/config"
	"task-microservice/db"
	_ "task-microservice/docs"
	"task-microservice/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Management API
// @version 1.0
// @description This is a sample API for managing tasks.
// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	config.LoadEnv()
	db.InitDB()
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(router)
	port := os.Getenv("PORT")
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
