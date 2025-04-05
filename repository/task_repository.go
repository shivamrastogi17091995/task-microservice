package repository

import (
	"strings"
	"task-microservice/db"
	"task-microservice/models"
	"time"

	"gorm.io/gorm"
)

func CreateTask(task models.Task) (models.Task, error) {
	task.ID = 0
	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt
	task.DeletedAt = gorm.DeletedAt{}
	task.Status = models.PENDING
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
	if task.Status == models.PENDING && strings.ToUpper(string(updatedData.Status)) == string(models.COMPLETED) {
		task.Status = models.COMPLETED
	}
	err = db.DB.Save(&task).Error
	return task, err
}

func DeleteTask(id uint) error {
	var task models.Task
	err := db.DB.First(&task, id).Error
	if err != nil {
		return err
	}
	return db.DB.Delete(&models.Task{}, id).Error
}

func GetAllTasks(status string, limit, offset int) ([]models.Task, error) {
	var tasks []models.Task
	query := db.DB.Order("id asc").Limit(limit).Offset(offset)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&tasks).Error
	return tasks, err
}
