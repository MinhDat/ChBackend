package seed

import (
	helper "ChGo/helpers"
	"ChGo/models"

	"syreclabs.com/go/faker"
)

var data = []string{"Sport", "Travel", "Music", "Game", "Photo", "Food"}
var data2 = []string{"Woman", "Man", "Child", "Animal", "Life", "Book"}

// Category seeding data
func Category() {
	db := helper.GetDB()
	var user models.User
	db.Where("username = ?", System_user).First(&user)

	createCategory(data, user.ID, models.CATEGORY_DEFAULT, models.IMAGE_PROFILE)
	createCategory(data2, user.ID, models.CATEGORY_TOPIC, models.IMAGE_PROFILE)
}

func createCategory(arr []string, oID int64, catType int8, imageType int8) {
	db := helper.GetDB()
	for _, value := range data {
		// Create category
		cat := models.Category{
			Name:    value,
			OwnerID: oID,
			Type:    catType,
		}
		db.Create(&cat)

		// Create Media for category
		media := models.Media{
			CorrelationID:  cat.ID,
			TableReference: "categories",
			Path:           faker.Avatar().Url("jpg", 500, 250),
			OwnerID:        cat.OwnerID,
			Type:           imageType,
			Extension:      "jpg",
		}
		db.Create(&media)
	}
}
