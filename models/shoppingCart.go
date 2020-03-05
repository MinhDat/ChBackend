package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// ShoppingCart "Object"
type ShoppingCart struct {
	common.Model
	OwnerID    uuid.UUID   `json:"owner_id" sql:"index"`
	ProductIDs common.JSON `sql:"type:json" json:"object"`
}
