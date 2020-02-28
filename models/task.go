package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Task "Object
type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool      `json:"completed"`
}

func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	task.ID, _ = uuid.NewV4()
	task.CreatedAt = time.Now()
	return nil
}

func (task *Task) BeforeUpdate(scope *gorm.Scope) error {
	task.UpdatedAt = time.Now()
	return nil
}
