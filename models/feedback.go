package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Feedback "Object"
type Feedback struct {
	common.Model
	ProductID  uuid.UUID `json:"product_id"`
	CustomerID string    `json:"customer_id"`
	Star       int       `json:"star"`
	Comment    string    `json:"comment"`
}
