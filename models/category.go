package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Category "Object
type Category struct {
	common.Model
	Name     string    `json:"name" binding:"required"`
	Image    string    `json:"image"`
	OwnerID  uuid.UUID `json:"owner_id"`
	ParentID uuid.UUID `json:"parent_id"`
}
