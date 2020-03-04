package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateRating(db *gorm.DB) {
	if !db.HasTable(&models.Rating{}) {
		err := db.CreateTable(&models.Rating{})
		if err != nil {
			log.Println("Rating table already exists")
		}
	}
}
