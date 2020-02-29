package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateProduct(db *gorm.DB) {
	if !db.HasTable(&models.Product{}) {
		err := db.CreateTable(&models.Product{})
		if err != nil {
			log.Println("Product table already exists")
		}
	}
	db.Model(&models.Product{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("category_id", "categories(uuid)", "RESTRICT", "RESTRICT")
}
