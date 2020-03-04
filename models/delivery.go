package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Delivery "Object
type Delivery struct {
	common.Model
	ShoppingCartID uuid.UUID `json:"shopping_cart_id"`
	Status         int       `json:"status"`
	Reason         string    `json:"reason"`
	Recipient      uuid.UUID `json:"recipient"`
	Sender         uuid.UUID `json:"sender"`
}
