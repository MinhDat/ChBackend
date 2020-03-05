package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateMedia() {
	db := db.GetDB()
	if !db.HasTable(&models.Media{}) {
		res := db.CreateTable(&models.Media{})
		if res != nil {
			log.Println("Media table has already created")
		}
	}
}
