package repository

import (
	"task-microservice/db"
	"task-microservice/models"
)

func CreateTask(task models.Task) (models.Task, error) {
	err := db.DB.Create(&task).Error
	return task, err
}
