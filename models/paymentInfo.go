package models

import (
	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// There is image in media table

// PaymentInfo "Object"
type PaymentInfo struct {
	common.Model
	PaymentID    uuid.UUID `json:"payment_id"`
	BrandName    string    `json:"brand_name"`
	Organization string    `json:"organization"`
}
