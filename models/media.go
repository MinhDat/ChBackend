package models

import (
	"ChGo/models/common"
)

const (
	IMAGE_THUMBNAIL = 0
	IMAGE_PROFILE   = 1
)

// Media "Object"
type Media struct {
	common.Model
	CorrelationID  int64   `json:"correlation_id"`
	TableReference string  `json:"table_reference" gorm:"type:varchar(100)"`
	Path           string  `json:"path" gorm:"type:varchar(255)"`
	Size           float32 `json:"size"`
	Type           int8    `json:"type"`
	OwnerID        int64   `json:"owner_id" sql:"index"`
	Extension      string  `json:"extension" gorm:"type:varchar(5)"`
}
