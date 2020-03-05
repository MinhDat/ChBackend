package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"log"
)

func MigrateOrder() {
	db := helper.GetDB()
	if !db.HasTable(&models.Order{}) {
		res := db.CreateTable(&models.Order{})
		if res != nil {
			log.Println("Order table has already created")
		}
	}
}
