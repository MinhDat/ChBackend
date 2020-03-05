package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateAuth() {
	db := db.GetDB()
	if !db.HasTable(&models.Auth{}) {
		err := db.CreateTable(&models.Auth{})
		if err != nil {
			log.Println("Auth table already exists")
		}
	}
}
