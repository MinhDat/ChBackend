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
	// OrderID uuid.UUID `json:"order_id" sql:"index"`
	Method string `json:"method"`
	Type   int    `json:"type"`
	Status int    `json:"status"`
}
