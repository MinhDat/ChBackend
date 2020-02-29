package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Product "Object
type ShoppingCart struct {
	common.Model
	UserID     uuid.UUID   `json:"user_id"`
	ProductIDs common.JSON `sql:"type:json" json:"object"`
}
