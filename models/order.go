package models

import (
	"ChGo/models/common"
)

// Order "Object"
type Order struct {
	common.Model
	OwnerID    int64       `json:"owner_id" sql:"index"`
	ProductIDs common.JSON `sql:"type:json" json:"object"`
}
