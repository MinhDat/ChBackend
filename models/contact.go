package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Contact "Object"
type Contact struct {
	common.Model
	OwnerID   uuid.UUID   `json:"owner_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Address   string      `json:"address"`
	Country   string      `json:"country"`
	National  string      `json:"national"`
	Social    common.JSON `json:"social"`
}
