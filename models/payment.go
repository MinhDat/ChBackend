package models

import (
	"ChGo/models/common"
)

const (
	CASHING = 0
	BANKING = 1
	VISA    = 2
)

// Payment "Object"
type Payment struct {
	common.Model
	// OrderID int `json:"order_id" sql:"index"`
	Method string `json:"method" gorm:"type:varchar(50)"`
	Type   int8   `json:"type"`
	Status int8   `json:"status"`
}
