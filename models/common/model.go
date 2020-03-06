package common

import (
	"time"

	"github.com/jinzhu/gorm"
)

// gorm.Model definition
type Model struct {
	ID        int64     `sql:"index" gorm:"AUTO_INCREMENT;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	model.CreatedAt = time.Now()
	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	model.UpdatedAt = time.Now()
	return nil
}
