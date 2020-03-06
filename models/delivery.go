package models

import (
	"ChGo/models/common"
)

// Delivery "Object
type Delivery struct {
	common.Model
	OrderID   int64  `json:"order_id" sql:"index"`
	Status    int8   `json:"status"`
	Reason    string `json:"reason" gorm:"type:varchar(255)"`
	Recipient int64  `json:"recipient" sql:"index"`
	Sender    int64  `json:"sender" sql:"index"`
}
