package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateCategory() {
	db := helper.GetDB()
	if !db.HasTable(&models.Category{}) {
		res := db.CreateTable(&models.Category{})
		if res != nil {
			log.Println("Category table has already created")
		}
	}
}
