package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateFeedback() {
	db := helper.GetDB()
	if !db.HasTable(&models.Feedback{}) {
		res := db.CreateTable(&models.Feedback{})
		if res != nil {
			log.Println("Feedback table has already created")
		}
	}
}
