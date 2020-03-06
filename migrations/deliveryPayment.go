package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateDeliveryPayment() {
	db := helper.GetDB()
	if !db.HasTable(&models.DeliveryPayment{}) {
		err := db.CreateTable(&models.DeliveryPayment{})
		if err != nil {
			log.Println("DeliveryPayment table already exists")
		}
	}
}
