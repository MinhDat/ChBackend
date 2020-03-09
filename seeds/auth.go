package seed

import (
	helper "ChGo/helpers"
	"ChGo/models"
)

var auths = []string{models.AUTH_ADMIN, models.AUTH_RETAILER, models.AUTH_CUSTOMER}

// Auth seeding data
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
