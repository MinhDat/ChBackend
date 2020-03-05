package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Meta "Object
type Meta struct {
	common.Model
	CorrelationID uuid.UUID   `json:"correlation_id" binding:"required"`
	Key           string      `json:"key" binding:"required"`
	Value         common.JSON `json:"value"`
}
