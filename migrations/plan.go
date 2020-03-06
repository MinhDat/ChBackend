package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigratePlan() {
	db := helper.GetDB()
	if !db.HasTable(&models.Plan{}) {
		err := db.CreateTable(&models.Plan{})
		if err != nil {
			log.Println("Plan table already exists")
		}
	}
}
