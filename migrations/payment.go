package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigratePayment() {
	db := helper.GetDB()
	if !db.HasTable(&models.Payment{}) {
		res := db.CreateTable(&models.Payment{})
		if res != nil {
			log.Println("Payment table has already created")
		}
	}
}
