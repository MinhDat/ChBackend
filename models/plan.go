package models

import (
	"ChGo/models/common"
)

// Plan "Object"
type Plan struct {
	common.Model
	Name          string      `json:"name" gorm:"type:varchar(50)"`
	Description   string      `json:"description" gorm:"type:varchar(255)"`
	Status        int8        `json:"status"`
	Specification common.JSON `json:"specification"`
}
