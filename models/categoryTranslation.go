package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Product "Object
type CategoryTranslation struct {
	common.Model
	CategoryID  uuid.UUID `json:"category_id"`
	Locale      string    `json:"name" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
}
