package models

import (
	"ChGo/models/common"
)

// UserPlan "Object"
type UserPlan struct {
	common.Model
	OwnerID int64 `json:"owner_id" sql:"index"`
	PlanID  int64 `json:"plan_id" sql:"index"`
}
