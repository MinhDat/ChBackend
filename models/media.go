package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

const (
	IMAGE_THUMBNAIL = 0
	IMAGE_PROFILE   = 1
)

// Media "Object"
type Media struct {
	common.Model
	CorrelationID  uuid.UUID `json:"correlation_id"`
	TableReference string    `json:"table_reference"`
	Path           string    `json:"path"`
	Size           float32   `json:"size"`
	Type           int       `json:"type"`
	OwnerID        uuid.UUID `json:"owner_id" sql:"index"`
	Extension      string    `json:"extension"`
}
