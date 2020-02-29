package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateFavorite(db *gorm.DB) {
	if !db.HasTable(&models.Favorite{}) {
		err := db.CreateTable(&models.Favorite{})
		if err != nil {
			log.Println("Favorite table already exists")
		}
	}
	db.Model(&models.Favorite{}).AddForeignKey("user_id", "users(uuid)", "RESTRICT", "RESTRICT")
}
