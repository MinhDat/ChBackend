package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// DeliveryPayment "Object
type DeliveryPayment struct {
	common.Model
	DeliveryID uuid.UUID `json:"delivery_id" sql:"index"`
	PaymentID  uuid.UUID `json:"payment_id" sql:"index"`
}
