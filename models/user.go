package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User "Object
type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `form:"username" json:"username" binding:"required"`
	Password  string    `form:"password" json:"password" binding:"required"`
	FirstName string    `json:"Firstname"`
	LastName  string    `json:"Lastname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	user.ID, _ = uuid.NewV4()
	user.CreatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	user.UpdatedAt = time.Now()
	return nil
}
