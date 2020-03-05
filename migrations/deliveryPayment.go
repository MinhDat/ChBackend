package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateDeliveryPayment() {
	db := db.GetDB()
	if !db.HasTable(&models.DeliveryPayment{}) {
		err := db.CreateTable(&models.DeliveryPayment{})
		if err != nil {
			log.Println("DeliveryPayment table already exists")
		}
	}
}
