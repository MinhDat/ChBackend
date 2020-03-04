package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateDelivery(db *gorm.DB) {
	if !db.HasTable(&models.Delivery{}) {
		err := db.CreateTable(&models.Delivery{})
		if err != nil {
			log.Println("Delivery table already exists")
		}
	}
}
