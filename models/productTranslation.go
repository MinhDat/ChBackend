package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// ProductTranslation "Object"
type ProductTranslation struct {
	common.Model
	ProductID   uuid.UUID `json:"product_id" sql:"index"`
	Locale      string    `json:"locale" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Origin      string    `json:"origin"`
}
