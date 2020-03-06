package seed

import (
	helper "ChGo/helpers"
	"ChGo/models"
)

var auths = []string{"Admin", "Retailer", "Customer"}

func Auth() {
	db := helper.GetDB()

	for _, value := range auths {
		auth := models.Auth{
			Name: value,
			Type: models.PERMISSION_NONE,
		}
		db.Create(&auth)
	}
}
