package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateAuth() {
	db := db.GetDB()
	if !db.HasTable(&models.Auth{}) {
		res := db.CreateTable(&models.Auth{})
		if res != nil {
			log.Println("Auth table has already created")
		}
	}
}
