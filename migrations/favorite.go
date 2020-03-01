package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateFavorite() {
	db := db.GetDB()
	if !db.HasTable(&models.Favorite{}) {
		res := db.CreateTable(&models.Favorite{})
		if res != nil {
			log.Println("Favorite table has already created")
		}
	}
}
