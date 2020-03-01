package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateAuth(db *gorm.DB) {
	if !db.HasTable(&models.Auth{}) {
		err := db.CreateTable(&models.Auth{})
		if err != nil {
			log.Println("Auth table already exists")
		}
	}
}
