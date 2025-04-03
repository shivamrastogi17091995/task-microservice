package services

import (
	"task-microservice/models"
	"task-microservice/repository"
)

func CreateTask(task models.Task) (models.Task, error) {
	return repository.CreateTask(task)
}

func GetTask(id uint) (models.Task, error) {
	return repository.GetTask(id)
}

func UpdateTask(id uint, task models.Task) (models.Task, error) {
	return repository.UpdateTask(id, task)
}

func DeleteTask(id uint) error {
	return repository.DeleteTask(id)
}
