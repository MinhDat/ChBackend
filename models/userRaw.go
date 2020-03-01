package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// UserRaw "Object"
type UserRaw struct {
	common.Model
	OwnerID    uuid.UUID   `json:"owner_id"`
	ProductIDs common.JSON `sql:"type:json" json:"product_ids"`
}
