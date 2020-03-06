package models

import (
	"ChGo/models/common"
)

// Meta "Object
type Meta struct {
	common.Model
	CorrelationID int64       `json:"correlation_id" binding:"required"`
	Key           string      `json:"key" binding:"required" gorm:"type:varchar(50)"`
	Value         common.JSON `json:"value"`
}
