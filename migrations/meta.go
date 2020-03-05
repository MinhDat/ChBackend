package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateMeta() {
	db := helper.GetDB()
	if !db.HasTable(&models.Meta{}) {
		res := db.CreateTable(&models.Meta{})
		if res != nil {
			log.Println("Meta table has already created")
		}
	}
}
