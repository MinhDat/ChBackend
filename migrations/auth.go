package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateAuth() {
	db := helper.GetDB()
	if !db.HasTable(&models.Auth{}) {
		err := db.CreateTable(&models.Auth{})
		if err != nil {
			log.Println("Auth table already exists")
		}
	}
}
