package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigratePlan() {
	db := db.GetDB()
	if !db.HasTable(&models.Plan{}) {
		err := db.CreateTable(&models.Plan{})
		if err != nil {
			log.Println("Plan table already exists")
		}
	}
}
