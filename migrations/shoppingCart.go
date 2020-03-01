package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateShoppingCart() {
	db := db.GetDB()
	if !db.HasTable(&models.ShoppingCart{}) {
		res := db.CreateTable(&models.ShoppingCart{})
		if res != nil {
			log.Println("ShoppingCart table has already created")
		}
	}
}
