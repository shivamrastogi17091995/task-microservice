package models

import "gorm.io/gorm"

type TaskStatus string

const (
	PENDING   TaskStatus = "PENDING"
	COMPLETED TaskStatus = "COMPLETED"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	Status      TaskStatus `gorm:"type:enum('PENDING', 'COMPLETED');default:'PENDING'"`
}
