package models

import (
	"ChGo/models/common"
)

// Favorite "Object"
type Favorite struct {
	common.Model
	OwnerID   int64 `json:"owner_id" sql:"index"`
	ProductID int64 `json:"product_id" sql:"index"`
}
