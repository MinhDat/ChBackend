package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Feedback "Object"
type Feedback struct {
	common.Model
	ProductID  uuid.UUID `json:"product_id" sql:"index"`
	CustomerID string    `json:"customer_id" sql:"index"`
	Star       int       `json:"star"`
	Comment    string    `json:"comment"`
}
