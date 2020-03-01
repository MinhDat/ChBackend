package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateCategoryTranslation() {
	db := db.GetDB()
	if !db.HasTable(&models.CategoryTranslation{}) {
		res := db.CreateTable(&models.CategoryTranslation{})
		if res != nil {
			log.Println("CategoryTranslation table has already created")
		}
	}
}
