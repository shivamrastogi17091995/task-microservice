package repository

import (
	"task-microservice/db"
	"task-microservice/models"
)

func CreateTask(task models.Task) (models.Task, error) {
	err := db.DB.Create(&task).Error
	return task, err
}

func GetTask(id uint) (models.Task, error) {
	var task models.Task
	err := db.DB.First(&task, id).Error
	return task, err
}

func UpdateTask(id uint, updatedData models.Task) (models.Task, error) {
	var task models.Task
	err := db.DB.First(&task, id).Error
	if err != nil {
		return task, err
	}
	task.Title = updatedData.Title
	task.Description = updatedData.Description
	task.Status = updatedData.Status
	err = db.DB.Save(&task).Error
	return task, err
}
