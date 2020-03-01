package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Product "Object
type Favorite struct {
	common.Model
	OwnerID    uuid.UUID `json:"owner_id"`
	ProductID uuid.UUID `json:"product_id"`
}
