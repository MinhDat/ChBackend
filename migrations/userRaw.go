package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateUserRaw() {
	db := helper.GetDB()
	if !db.HasTable(&models.UserRaw{}) {
		res := db.CreateTable(&models.UserRaw{})
		if res != nil {
			log.Println("UserRaw table has already created")
		}
	}
}
