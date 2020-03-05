package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// CategoryTranslation "Object"
type CategoryTranslation struct {
	common.Model
	CategoryID  uuid.UUID `json:"category_id" sql:"index"`
	Locale      string    `json:"locale" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
}
