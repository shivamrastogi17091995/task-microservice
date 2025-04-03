package services

import (
	"task-microservice/models"
	"task-microservice/repository"
)

func CreateTask(task models.Task) (models.Task, error) {
	return repository.CreateTask(task)
}

func UpdateTask(id uint, task models.Task) (models.Task, error) {
	return repository.UpdateTask(id, task)
}
