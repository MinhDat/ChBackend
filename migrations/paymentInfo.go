package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigratePaymentInfo() {
	db := db.GetDB()
	if !db.HasTable(&models.PaymentInfo{}) {
		res := db.CreateTable(&models.PaymentInfo{})
		if res != nil {
			log.Println("PaymentInfo table has already created")
		}
	}
}
