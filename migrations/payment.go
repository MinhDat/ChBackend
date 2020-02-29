package migrate

import (
	"ChGo/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigratePayment(db *gorm.DB) {
	if !db.HasTable(&models.Payment{}) {
		err := db.CreateTable(&models.Payment{})
		if err != nil {
			log.Println("Payment table already exists")
		}
	}
	db.Model(&models.Payment{}).AddForeignKey("shopping_cart_id", "shopping_carts(uuid)", "RESTRICT", "RESTRICT")
}
