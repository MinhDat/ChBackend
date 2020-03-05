package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateUserRaw() {
	db := db.GetDB()
	if !db.HasTable(&models.UserRaw{}) {
		res := db.CreateTable(&models.UserRaw{})
		if res != nil {
			log.Println("UserRaw table has already created")
		}
	}
}
