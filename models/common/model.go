package common

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// gorm.Model definition
type Model struct {
	ID        int       `json:"id" sql:"index" gorm:"AUTO_INCREMENT"`
	UUID      uuid.UUID `json:"uuid" sql:"index" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
