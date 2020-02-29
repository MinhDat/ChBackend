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
	migrate.MigrateUser(db)
	migrate.MigrateCategory(db)
	migrate.MigrateCategoryTranslation(db)
	migrate.MigrateProduct(db)
	migrate.MigrateProductTranslation(db)
	migrate.MigrateShoppingCart(db)
	migrate.MigratePayment(db)
	migrate.MigrateFavorite(db)

	db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.CategoryTranslation{},
		&models.Product{},
		&models.ProductTranslation{},
		&models.ShoppingCart{},
		&models.Payment{},
		&models.Favorite{},
	)
	db.Close()
}
