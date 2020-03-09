package models

import (
	"ChGo/models/common"
)

// UserAuth "Object"
type UserAuth struct {
	common.Model
	OwnerID int64  `json:"owner_id" sql:"index"`
	AuthID  string `json:"auth_id" sql:"index" gorm:"type:varchar(50)"`
}
