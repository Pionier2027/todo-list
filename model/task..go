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

// dbからtask一覧を取得
func GetTasks() ([]Task, error) {
	var tasks []Task

	err := db.Find(&tasks).Error

	return tasks, err
}

// dbにtaskを追加
func AddTask(name string) (*Task, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	task := Task{
		ID:       id,
		Name:     name,
		Finished: false,
	}

	// taskをDBに追加
	err = db.Create(&task).Error

	return &task, err
}
