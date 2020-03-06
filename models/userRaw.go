package models

import (
	"ChGo/models/common"
)

// UserRaw "Object"
type UserRaw struct {
	common.Model
	Email   string `json:"email" gorm:"type:varchar(100)"`
	Phone   string `json:"phone" gorm:"type:varchar(20)"`
	Address string `json:"address" gorm:"type:varchar(255)"`
}
