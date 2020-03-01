package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigratePayment() {
	db := db.GetDB()
	if !db.HasTable(&models.Payment{}) {
		res := db.CreateTable(&models.Payment{})
		if res != nil {
			log.Println("Payment table has already created")
		}
	}
}
