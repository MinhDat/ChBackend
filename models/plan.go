package models

import (
	"ChGo/models/common"
)

// Plan "Object"
type Plan struct {
	common.Model
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Status        int         `json:"status"`
	Specification common.JSON `json:"specification"`
}
