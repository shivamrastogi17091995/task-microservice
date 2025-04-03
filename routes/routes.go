package routes

import (
	"task-microservice/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
}
