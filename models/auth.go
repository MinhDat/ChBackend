package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Auth "Object
type Auth struct {
	common.Model
	OwnerID   uuid.UUID `json:"owner_id"`
	Name      string    `json:"name" binding:"required"`
	Reference string    `json:"reference"`
	Type      int       `json:"type"`
}
