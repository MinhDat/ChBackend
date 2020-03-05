package models

import (
	"ChGo/models/common"
)

// UserRaw "Object"
type UserRaw struct {
	common.Model
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
