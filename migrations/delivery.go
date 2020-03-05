package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateDelivery() {
	db := db.GetDB()
	if !db.HasTable(&models.Delivery{}) {
		err := db.CreateTable(&models.Delivery{})
		if err != nil {
			log.Println("Delivery table already exists")
		}
	}
}
