package models

import (
	"ChGo/models/common"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// There is image in media table

// User "Object"
type User struct {
	common.Model
	Username string `form:"username" json:"username" binding:"required" gorm:"type:varchar(50)"`
	Password string `form:"password" json:"password" binding:"required"`
}

// AfterCreate ..
func (user *User) AfterCreate(scope *gorm.Scope) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	scope.DB().Model(user).Update("password", hash)
	return nil
}
