package models

import (
	"ChGo/models/common"
)

// DeliveryPayment "Object
type DeliveryPayment struct {
	common.Model
	DeliveryID int64 `json:"delivery_id" sql:"index"`
	PaymentID  int64 `json:"payment_id" sql:"index"`
}
