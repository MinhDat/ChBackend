package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateUser() {
	db := db.GetDB()
	if !db.HasTable(&models.User{}) {
		res := db.CreateTable(&models.User{})
		if res != nil {
			log.Println("User table has already created")
		}
	}
}
