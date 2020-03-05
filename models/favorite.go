package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Favorite "Object"
type Favorite struct {
	common.Model
	OwnerID   uuid.UUID `json:"owner_id" sql:"index"`
	ProductID uuid.UUID `json:"product_id" sql:"index"`
}
