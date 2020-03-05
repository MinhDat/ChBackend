package main

import (
	"ChGo/db"
	"ChGo/models"

	migrate "ChGo/migrations"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	db.Init()
	db := db.GetDB()

	// Migrate database
	migrate.MigrateAuth()
	migrate.MigrateCategory()
	migrate.MigrateCategoryTranslation()
	migrate.MigrateContact()
	migrate.MigrateDelivery()
	migrate.MigrateDeliveryPayment()
	migrate.MigrateFavorite()
	migrate.MigrateFeedback()
	migrate.MigrateMedia()
	migrate.MigrateMeta()
	migrate.MigratePayment()
	migrate.MigratePaymentInfo()
	migrate.MigratePermission()
	migrate.MigratePlan()
	migrate.MigrateProduct()
	migrate.MigrateProductTranslation()
	migrate.MigrateShoppingCart()
	migrate.MigrateUser()
	migrate.MigrateUserPlan()
	migrate.MigrateUserRaw()
	migrate.MigrateAssociation()

	db.AutoMigrate(
		&models.Auth{},
		&models.Category{},
		&models.CategoryTranslation{},
		&models.Contact{},
		&models.Delivery{},
		&models.DeliveryPayment{},
		&models.Favorite{},
		&models.Media{},
		&models.Meta{},
		&models.Payment{},
		&models.PaymentInfo{},
		&models.Product{},
		&models.ProductTranslation{},
		&models.Feedback{},
		&models.ShoppingCart{},
		&models.User{},
		&models.UserPlan{},
		&models.UserRaw{},
	)
	db.Close()
}
