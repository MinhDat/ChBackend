package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
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
