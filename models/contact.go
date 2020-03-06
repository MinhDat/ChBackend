package models

import (
	"ChGo/models/common"
	"time"
)

// Contact "Object
type Contact struct {
	ID        int64       `sql:"index" gorm:"primary_key"`
	FirstName string      `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string      `json:"last_name" gorm:"type:varchar(50)"`
	Email     string      `json:"email" gorm:"type:varchar(100);unique_index"`
	Phone     string      `json:"phone" gorm:"type:varchar(20);unique_index"`
	Address   string      `json:"address" gorm:"type:varchar(255)"`
	Country   string      `json:"country" gorm:"type:varchar(50)"`
	National  string      `json:"national" gorm:"type:varchar(50)"`
	Social    common.JSON `json:"social"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
