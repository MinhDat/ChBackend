package models

import (
	"ChGo/models/common"
)

// UserAuth "Object"
type UserAuth struct {
	common.Model
	OwnerID int64 `json:"owner_id" sql:"index"`
	AuthID  int64 `json:"plan_id" sql:"index"`
}
