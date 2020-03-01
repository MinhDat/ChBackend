package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateCategory(db *gorm.DB) {
	if !db.HasTable(&models.Category{}) {
		err := db.CreateTable(&models.Category{})
		if err != nil {
			log.Println("Category table already exists")
		}
	}
}
