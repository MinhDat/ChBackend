package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateShoppingCart(db *gorm.DB) {
	if !db.HasTable(&models.ShoppingCart{}) {
		err := db.CreateTable(&models.ShoppingCart{})
		if err != nil {
			log.Println("ShoppingCart table already exists")
		}
	}
}
