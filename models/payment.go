package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Payment "Object"
type Payment struct {
	common.Model
	ShoppingCartID uuid.UUID `json:"shopping_cart_id"`
	Method         string    `json:"method"`
	Status         int
}
