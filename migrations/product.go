package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateProduct() {
	db := helper.GetDB()
	if !db.HasTable(&models.Product{}) {
		res := db.CreateTable(&models.Product{})
		if res != nil {
			log.Println("Product table has already created")
		}
	}
}
