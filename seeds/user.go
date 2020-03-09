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

const (
	System_user = "admin"
	System_pass = "admin"
)

// User seeding data
func User() {
	var db = helper.GetDB()
	user := models.User{
		Username: System_user,
		Password: System_pass,
	}

	db.Create(&user)

	helper.Write("user.csv", [][]string{{"usename", "password"}})
	helper.Write("user.csv", [][]string{{System_user, System_pass}})

	// Make admin
	createUser(TOTAL_ADMIN, models.AUTH_ADMIN)
	// Make retailer
	createUser(TOTAL_RETAILER, models.AUTH_RETAILER)
	// Make customer
	createUser(TOTAL_CUSTOMER, models.AUTH_CUSTOMER)

}

func createUser(amount int, auth string) {
	var db = helper.GetDB()
	for i := 0; i < amount; i++ {
		// Save user data seeding into csv file
		user := models.User{
			Username: faker.Internet().UserName(),
			Password: faker.Internet().Password(8, 8),
		}
		helper.Write("user.csv", [][]string{{user.Username, user.Password}})

		// Create User
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

		// Create User Media
		media := models.Media{
			CorrelationID:  user.ID,
			TableReference: "user",
			Path:           faker.Avatar().Url("jpg", 200, 200),
			OwnerID:        user.ID,
			Type:           models.IMAGE_THUMBNAIL,
			Extension:      "jpg",
		}
		db.Create(&media)

		// Create User-Autheration
		userAuth := models.UserAuth{
			OwnerID: user.ID,
			AuthID:  auth,
		}

		db.Create(&userAuth)
	}
}
