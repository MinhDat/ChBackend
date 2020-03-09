package models

import (
	"ChGo/models/common"
)

const (
	PERMISSION_NONE     = 0
	PERMISSION_REQUIRED = 1
)

const (
	AUTH_ADMIN    = "Admin"
	AUTH_RETAILER = "Retailer"
	AUTH_CUSTOMER = "Customer"
)

// Auth "Object
type Auth struct {
	common.Model
	Name         string `json:"name" binding:"required" gorm:"type:varchar(50);unique;not null"`
	PermissionID int64  `json:"owner_id" sql:"index"`
	Type         int8   `json:"type"`
}
