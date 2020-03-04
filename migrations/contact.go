package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateContact(db *gorm.DB) {
	if !db.HasTable(&models.Contact{}) {
		err := db.CreateTable(&models.Contact{})
		if err != nil {
			log.Println("Contact table already exists")
		}
	}
}
