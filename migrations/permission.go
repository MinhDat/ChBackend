package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigratePermission(db *gorm.DB) {
	if !db.HasTable(&models.Permission{}) {
		err := db.CreateTable(&models.Permission{})
		if err != nil {
			log.Println("Permission table already exists")
		}
	}
}
