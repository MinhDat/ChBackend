package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

const (
	thumbnail = 0
	image     = 1
)

// Media "Object"
type Media struct {
	common.Model
	CorrelationID uuid.UUID `json:"correlation_id"`
	Path          string    `json:"path"`
	Size          float32   `json:"size"`
	Type          int       `json:"type"`
	OwnerID       uuid.UUID `json:"owner_id"`
	Extension     string    `json:"extension"`
}
