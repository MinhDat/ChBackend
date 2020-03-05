package migrate

import (
	"ChGo/db"
	"ChGo/models"
	"log"
)

func MigrateUserPlan() {
	db := db.GetDB()
	if !db.HasTable(&models.UserPlan{}) {
		res := db.CreateTable(&models.UserPlan{})
		if res != nil {
			log.Println("UserPlan table has already created")
		}
	}
}
