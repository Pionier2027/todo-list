package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Task struct {
	ID       uuid.UUID
	Name     string
	Finished bool
}

// dbからtask一覧を取得する関数
func GetTasks() ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}
