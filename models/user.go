package models

import (
	"ChGo/models/common"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User "Object
type User struct {
	common.Model
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	user.UUID, _ = uuid.NewV4()
	user.CreatedAt = time.Now()
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(hash)
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	user.UpdatedAt = time.Now()
	return nil
}
