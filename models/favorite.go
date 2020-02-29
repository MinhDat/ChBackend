package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Product "Object
type Favorite struct {
	common.Model
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
}
