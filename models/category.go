package models

import (
	"ChGo/models/common"
)

// Constant for category type
const (
	CATEGORY_DEFAULT  = 0
	CATEGORY_TOPIC    = 1
	CATEGORY_HOT_PICK = 2
)

// There is image in media table

// Category "Object"
type Category struct {
	common.Model
	Name     string `json:"name" binding:"required" gorm:"type:varchar(50)"`
	OwnerID  int64  `json:"owner_id" sql:"index"`
	ParentID int64  `json:"parent_id" sql:"index"`
	Type     int8   `json:"type"`
}
