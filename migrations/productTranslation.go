package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateProductTranslation() {
	db := helper.GetDB()
	if !db.HasTable(&models.ProductTranslation{}) {
		res := db.CreateTable(&models.ProductTranslation{})
		if res != nil {
			log.Println("ProductTranslation table has already created")
		}
	}
}
