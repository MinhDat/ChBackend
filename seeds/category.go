package seed

import (
	helper "ChGo/helpers"
	"ChGo/models"
	"math/rand"
	"time"

	"syreclabs.com/go/faker"
)

const (
	MAX_ADMIN    = 5
	MAX_RETAILER = 5
	MAX_CUSTOMER = 5
)

var data = []string{"Sport", "Travel", "Music", "Game", "Photo", "Food"}
var data2 = []string{"Woman", "Man", "Child", "Animal", "Life", "Book"}

func Category() {
	db := helper.GetDB()
	var users []models.User
	db.Find(&users)

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(users)

	for _, value := range data {
		oID := rand.Intn(max - min)
		cat := models.Category{
			Name:    value,
			OwnerID: users[oID].UUID,
			Type:    models.CATEGORY_DEFAULT,
		}
		db.Create(&cat)

		media := models.Media{
			CorrelationID:  cat.UUID,
			TableReference: "categories",
			Path:           faker.Avatar().Url("jpg", 500, 250),
			OwnerID:        cat.OwnerID,
			Type:           models.IMAGE_PROFILE,
			Extension:      "jpg",
		}
		db.Create(&media)
	}

	for _, value := range data2 {
		oID := rand.Intn(max - min)
		cat := models.Category{
			Name:    value,
			OwnerID: users[oID].UUID,
			Type:    models.CATEGORY_TOPIC,
		}
		db.Create(&cat)

		media := models.Media{
			CorrelationID:  cat.UUID,
			TableReference: "categories",
			Path:           faker.Avatar().Url("jpg", 500, 250),
			OwnerID:        cat.OwnerID,
			Type:           models.IMAGE_PROFILE,
			Extension:      "jpg",
		}
		db.Create(&media)
	}
}
