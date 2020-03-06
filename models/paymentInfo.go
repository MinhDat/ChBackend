package models

import (
	"ChGo/models/common"
)

// There is image in media table

// PaymentInfo "Object"
type PaymentInfo struct {
	common.Model
	PaymentID    int64  `json:"payment_id" sql:"index"`
	BrandName    string `json:"brand_name" gorm:"type:varchar(100)"`
	Organization string `json:"organization" gorm:"type:varchar(100)"`
}
