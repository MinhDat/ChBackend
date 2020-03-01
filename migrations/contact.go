package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateContact() {
	db := db.GetDB()
	if !db.HasTable(&models.Contact{}) {
		res := db.CreateTable(&models.Contact{})
		if res != nil {
			log.Println("Contact table has already created")
		}
	}
}
