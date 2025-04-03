package services

import (
	"task-microservice/models"
	"task-microservice/repository"
)

func CreateTask(task models.Task) (models.Task, error) {
	return repository.CreateTask(task)
}
