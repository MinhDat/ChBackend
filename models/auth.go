package models

import (
	"ChGo/models/common"
)

const (
	PERMISSION_NONE     = 0
	PERMISSION_REQUIRED = 1
)

// Auth "Object
type Auth struct {
	common.Model
	Name         string `json:"name" binding:"required" gorm:"type:varchar(50)"`
	PermissionID int64  `json:"owner_id" sql:"index"`
	Type         int8   `json:"type"`
}
