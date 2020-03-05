package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
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
