package models

import (
	"ChGo/models/common"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User "Object"
type User struct {
	common.Model
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// AfterCreate ..
func (user *User) AfterCreate(scope *gorm.Scope) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	scope.DB().Model(user).Update("password", hash)
	return nil
}
