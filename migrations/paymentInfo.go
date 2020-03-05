package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigratePaymentInfo() {
	db := helper.GetDB()
	if !db.HasTable(&models.PaymentInfo{}) {
		res := db.CreateTable(&models.PaymentInfo{})
		if res != nil {
			log.Println("PaymentInfo table has already created")
		}
	}
}
