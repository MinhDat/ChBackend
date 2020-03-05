package models

import (
	"ChGo/models/common"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Contact "Object
type Contact struct {
	ID        int         `json:"-" sql:"index" gorm:"AUTO_INCREMENT"`
	UUID      uuid.UUID   `json:"uuid" sql:"index" gorm:"primary_key"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Address   string      `json:"address"`
	Country   string      `json:"country"`
	National  string      `json:"national"`
	Social    common.JSON `json:"social"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
