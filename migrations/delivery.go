package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateDelivery() {
	db := db.GetDB()
	if !db.HasTable(&models.Delivery{}) {
		res := db.CreateTable(&models.Delivery{})
		if res != nil {
			log.Println("Delivery table has already created")
		}
	}
}
