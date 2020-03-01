package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateProductTranslation(db *gorm.DB) {
	if !db.HasTable(&models.ProductTranslation{}) {
		err := db.CreateTable(&models.ProductTranslation{})
		if err != nil {
			log.Println("ProductTranslation table already exists")
		}
	}
}
