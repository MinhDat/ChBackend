package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateUserPlan() {
	db := helper.GetDB()
	if !db.HasTable(&models.UserPlan{}) {
		res := db.CreateTable(&models.UserPlan{})
		if res != nil {
			log.Println("UserPlan table has already created")
		}
	}
}
