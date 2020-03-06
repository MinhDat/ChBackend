package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigratePermission() {
	db := helper.GetDB()
	if !db.HasTable(&models.Permission{}) {
		err := db.CreateTable(&models.Permission{})
		if err != nil {
			log.Println("Permission table already exists")
		}
	}
}
