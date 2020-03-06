package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateUserAuth() {
	db := helper.GetDB()
	if !db.HasTable(&models.UserAuth{}) {
		res := db.CreateTable(&models.UserAuth{})
		if res != nil {
			log.Println("UserAuth table has already created")
		}
	}
}
