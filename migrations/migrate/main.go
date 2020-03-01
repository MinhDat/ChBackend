package main

import (
	"ChGo/models"
	"fmt"
	"log"
	"os"

	migrate "ChGo/migrations"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	user := getEnv("PG_USER", "hugo")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "8080")
	database := getEnv("PG_DB", "tasks")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	// Migrate database
	migrate.MigrateAuth(db)
	migrate.MigrateCategory(db)
	migrate.MigrateCategoryTranslation(db)
	migrate.MigrateContact(db)
	migrate.MigrateDelivery(db)
	migrate.MigrateFavorite(db)
	migrate.MigratePayment(db)
	migrate.MigratePermission(db)
	migrate.MigrateProduct(db)
	migrate.MigrateProductTranslation(db)
	migrate.MigrateRating(db)
	migrate.MigrateShoppingCart(db)
	migrate.MigrateUser(db)
	migrate.MigrateAssociation(db)

	db.AutoMigrate(
		&models.Auth{},
		&models.Category{},
		&models.CategoryTranslation{},
		&models.Contact{},
		&models.Delivery{},
		&models.Favorite{},
		&models.Payment{},
		&models.Product{},
		&models.ProductTranslation{},
		&models.Rating{},
		&models.ShoppingCart{},
		&models.User{},
	)
	db.Close()
}
