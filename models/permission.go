package models

import (
	"ChGo/models/common"
)

// Permission "Object
type Permission struct {
	common.Model
	AuthID int64  `json:"auth_id" sql:"index"`
	Method string `json:"name" gorm:"type:varchar(6)"`
	Path   string `json:"path" gorm:"type:varchar(255)"`
}
