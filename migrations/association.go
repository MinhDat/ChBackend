package migrate

import (
	"ChGo/db"
	"ChGo/models"
)

// MigrateAssociation ..
func MigrateAssociation() {
	db := db.GetDB()
	db.Model(&models.Auth{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Category{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.CategoryTranslation{}).AddForeignKey("category_id", "categories(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Contact{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Delivery{}).AddForeignKey("shopping_cart_id", "shopping_carts(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Favorite{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Payment{}).AddForeignKey("shopping_cart_id", "shopping_carts(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Permission{}).AddForeignKey("auth_id", "auths(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("category_id", "categories(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.ProductTranslation{}).AddForeignKey("product_id", "products(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Feedback{}).AddForeignKey("product_id", "products(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.ShoppingCart{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
}
