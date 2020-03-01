package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateProduct() {
	db := db.GetDB()
	if !db.HasTable(&models.Product{}) {
		res := db.CreateTable(&models.Product{})
		if res != nil {
			log.Println("Product table has already created")
		}
	}
}
