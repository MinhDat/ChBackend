package models

import (
	"time"

	"ChGo/models/common"
)

// There is image in media table

// Product "Object"
type Product struct {
	common.Model
	Name           string    `json:"name" binding:"required" gorm:"type:varchar(50)"`
	OwnerID        int64     `json:"owner_id" sql:"index"`
	CategoryID     int64     `json:"category_id" sql:"index"`
	Producer       string    `json:"producer" gorm:"type:varchar(50)"`
	Material       string    `json:"material" gorm:"type:varchar(50)"`
	ProductionDate time.Time `json:"production_date"`
	Price          float32   `json:"price"`
	SaleOff        float32   `json:"sale_off"`
}
