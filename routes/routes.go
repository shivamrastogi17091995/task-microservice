package routes

import (
	"task-microservice/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", handlers.CreateTask)
}
