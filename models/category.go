package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Constant for category type
const (
	CATEGORY_DEFAULT  = 0
	CATEGORY_TOPIC    = 1
	CATEGORY_HOT_PICK = 2
)

// There is image in media table

// Category "Object"
type Category struct {
	common.Model
	Name     string    `json:"name" binding:"required"`
	OwnerID  uuid.UUID `json:"owner_id" sql:"index"`
	ParentID uuid.UUID `json:"parent_id" sql:"index"`
	Type     int       `json:"type"`
}
