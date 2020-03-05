package common

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// gorm.Model definition
type Model struct {
	ID        int       `json:"-" sql:"index" gorm:"AUTO_INCREMENT"`
	UUID      uuid.UUID `json:"uuid" sql:"index" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	if model.UUID == uuid.Nil {
		model.UUID, _ = uuid.NewV4()
	}
	model.CreatedAt = time.Now()
	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	model.UpdatedAt = time.Now()
	return nil
}
