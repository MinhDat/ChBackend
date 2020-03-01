package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateFeedback() {
	db := db.GetDB()
	if !db.HasTable(&models.Feedback{}) {
		res := db.CreateTable(&models.Feedback{})
		if res != nil {
			log.Println("Feedback table has already created")
		}
	}
}
