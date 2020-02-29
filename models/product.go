package models

import (
	"time"

	"ChGo/models/common"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Product "Object
type Product struct {
	common.Model
	Name    string    `json:"name" binding:"required"`
	OwnerID uuid.UUID `json:"owner_id"`
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	product.UUID, _ = uuid.NewV4()
	product.CreatedAt = time.Now()
	return nil
}

func (product *Product) BeforeUpdate(scope *gorm.Scope) error {
	product.UpdatedAt = time.Now()
	return nil
}
