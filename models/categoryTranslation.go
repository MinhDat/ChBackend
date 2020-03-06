package models

import (
	"ChGo/models/common"
)

// CategoryTranslation "Object"
type CategoryTranslation struct {
	common.Model
	CategoryID  int64  `json:"category_id" sql:"index"`
	Locale      string `json:"locale" binding:"required" gorm:"type:varchar(5)"`
	Name        string `json:"name" binding:"required" gorm:"type:varchar(50)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}
