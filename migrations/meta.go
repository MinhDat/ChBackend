package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateMeta() {
	db := db.GetDB()
	if !db.HasTable(&models.Meta{}) {
		res := db.CreateTable(&models.Meta{})
		if res != nil {
			log.Println("Meta table has already created")
		}
	}
}
