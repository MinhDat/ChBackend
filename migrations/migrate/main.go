package main

import (
	helper "ChGo/helpers"
	"ChGo/models"

	migrate "ChGo/migrations"
)

func main() {
	helper.Init()
	db := helper.GetDB()

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
	migrate.MigrateOrder()
	migrate.MigrateUser()
	migrate.MigrateUserAuth()
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
		&models.Order{},
		&models.User{},
		&models.UserAuth{},
		&models.UserPlan{},
		&models.UserRaw{},
	)
	db.Close()
}
