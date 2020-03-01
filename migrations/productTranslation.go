package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateProductTranslation() {
	db := db.GetDB()
	if !db.HasTable(&models.ProductTranslation{}) {
		res := db.CreateTable(&models.ProductTranslation{})
		if res != nil {
			log.Println("ProductTranslation table has already created")
		}
	}
}
