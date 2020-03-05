package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Permission "Object
type Permission struct {
	common.Model
	AuthID uuid.UUID `json:"auth_id" sql:"index"`
	Method string    `json:"name"`
	Path   string    `json:"path"`
}
