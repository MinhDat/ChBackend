package models

import (
	"ChGo/models/common"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Feedback "Object"
type Feedback struct {
	common.Model
	ProductID  int64     `json:"product_id" sql:"index"`
	CustomerID uuid.UUID `json:"customer_id" sql:"index"`
	Star       float32   `json:"star"`
	Comment    string    `json:"comment" gorm:"type:varchar(255)"`
}

// AfterCreate ..
func (feedback *Feedback) AfterCreate(scope *gorm.Scope) error {
	customerID, _ := uuid.NewV4()
	scope.DB().Model(feedback).Update("customer_id", customerID)
	return nil
}
