package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateCategoryTranslation(db *gorm.DB) {
	if !db.HasTable(&models.CategoryTranslation{}) {
		err := db.CreateTable(&models.CategoryTranslation{})
		if err != nil {
			log.Println("CategoryTranslation table already exists")
		}
	}
}
