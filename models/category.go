package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Constant for category type
const (
	DEFAULT  = 0
	TOPIC    = 1
	HOT_PICK = 2
)

// There is image in media table

// Category "Object"
type Category struct {
	common.Model
	Name     string    `json:"name" binding:"required"`
	Image    string    `json:"image"`
	OwnerID  uuid.UUID `json:"owner_id" sql:"index"`
	ParentID uuid.UUID `json:"parent_id"`
	Type     int       `json:"type"`
}
