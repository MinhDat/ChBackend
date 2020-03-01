package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateCategory() {
	db := db.GetDB()
	if !db.HasTable(&models.Category{}) {
		res := db.CreateTable(&models.Category{})
		if res != nil {
			log.Println("Category table has already created")
		}
	}
}
