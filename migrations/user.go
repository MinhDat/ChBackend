package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateUser() {
	db := helper.GetDB()
	if !db.HasTable(&models.User{}) {
		res := db.CreateTable(&models.User{})
		if res != nil {
			log.Println("User table has already created")
		}
	}
}
