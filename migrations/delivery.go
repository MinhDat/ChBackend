package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateDelivery() {
	db := helper.GetDB()
	if !db.HasTable(&models.Delivery{}) {
		err := db.CreateTable(&models.Delivery{})
		if err != nil {
			log.Println("Delivery table already exists")
		}
	}
}
