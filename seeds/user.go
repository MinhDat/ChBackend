package seed

import (
	helper "ChGo/helpers"
	"ChGo/models"

	"syreclabs.com/go/faker"
)

const (
	TOTAL_ADMIN    = 5
	TOTAL_RETAILER = 5
	TOTAL_CUSTOMER = 5
)

func User() {
	var db = helper.GetDB()
	user := models.User{
		Username: "admin",
		Password: "admin",
	}

	db.Create(&user)

	helper.Write("user.csv", [][]string{{"usename", "password"}})
	helper.Write("user.csv", [][]string{{"admin", "admin"}})

	// Make admin
	for i := 0; i < TOTAL_ADMIN; i++ {
		// username := faker.Internet().UserName()
		// if len(username) > 50 {
		// 	username = username[0:49]
		// }
		user := models.User{
			Username: faker.Internet().UserName(),
			Password: faker.Internet().Password(8, 8),
		}
		helper.Write("user.csv", [][]string{{user.Username, user.Password}})

		db.Create(&user)
		contact := models.Contact{
			ID:        user.ID,
			FirstName: faker.Name().FirstName(),
			LastName:  faker.Name().LastName(),
			Email:     faker.Internet().Email(),
			Phone:     faker.PhoneNumber().CellPhone(),
			Address:   faker.Address().String(),
			Country:   faker.Address().Country(),
			National:  faker.Address().Country(),
		}
		db.Create(&contact)

		media := models.Media{
			CorrelationID:  user.ID,
			TableReference: "user",
			Path:           faker.Avatar().Url("jpg", 200, 200),
			OwnerID:        user.ID,
			Type:           models.IMAGE_THUMBNAIL,
			Extension:      "jpg",
		}
		db.Create(&media)
	}
}
