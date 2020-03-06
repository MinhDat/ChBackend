package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateContact() {
	db := helper.GetDB()
	if !db.HasTable(&models.Contact{}) {
		err := db.CreateTable(&models.Contact{})
		if err != nil {
			log.Println("Contact table already exists")
		}
	}
}
