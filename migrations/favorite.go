package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateFavorite() {
	db := helper.GetDB()
	if !db.HasTable(&models.Favorite{}) {
		res := db.CreateTable(&models.Favorite{})
		if res != nil {
			log.Println("Favorite table has already created")
		}
	}
}
