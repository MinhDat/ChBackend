package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigratePermission() {
	db := db.GetDB()
	if !db.HasTable(&models.Permission{}) {
		res := db.CreateTable(&models.Permission{})
		if res != nil {
			log.Println("Permission table has already created")
		}
	}
}
