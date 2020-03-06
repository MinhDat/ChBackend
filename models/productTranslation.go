package models

import (
	"ChGo/models/common"
)

// ProductTranslation "Object"
type ProductTranslation struct {
	common.Model
	ProductID   int64  `json:"product_id" sql:"index"`
	Locale      string `json:"locale" binding:"required" gorm:"type:varchar(5)"`
	Name        string `json:"name" binding:"required" gorm:"type:varchar(50)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Origin      string `json:"origin" gorm:"type:varchar(100)"`
}
