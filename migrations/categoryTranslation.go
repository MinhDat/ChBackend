package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateCategoryTranslation() {
	db := helper.GetDB()
	if !db.HasTable(&models.CategoryTranslation{}) {
		res := db.CreateTable(&models.CategoryTranslation{})
		if res != nil {
			log.Println("CategoryTranslation table has already created")
		}
	}
}
