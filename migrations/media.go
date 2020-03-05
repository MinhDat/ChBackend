package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateMedia() {
	db := helper.GetDB()
	if !db.HasTable(&models.Media{}) {
		res := db.CreateTable(&models.Media{})
		if res != nil {
			log.Println("Media table has already created")
		}
	}
}
