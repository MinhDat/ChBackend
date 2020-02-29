package models

import (
	"time"

	"ChGo/models/common"

	uuid "github.com/satori/go.uuid"
)

// Product "Object
type Product struct {
	common.Model
	Name           string    `json:"name" binding:"required"`
	OwnerID        uuid.UUID `json:"owner_id"`
	CategoryID     uuid.UUID `json:"category_id"`
	Image          string    `json:"image"`
	Thumbnail      string    `json:"thumbnail"`
	Producer       string    `json:"producer"`
	Material       string    `json:"material"`
	ProductionDate time.Time `json:"production_date"`
	price          float32   `json:"price"`
}
