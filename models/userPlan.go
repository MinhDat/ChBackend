package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// UserPlan "Object"
type UserPlan struct {
	common.Model
	OwnerID uuid.UUID `json:"owner_id" sql:"index"`
	PlanID  uuid.UUID `json:"plan_id" sql:"index"`
}
