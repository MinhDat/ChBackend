package migrate

import (
	"ChGo/db"
	"ChGo/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateAssociation() {
	db := db.GetDB()
	db.Model(&models.Auth{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Category{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.CategoryTranslation{}).AddForeignKey("category_id", "categories(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Contact{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Delivery{}).AddForeignKey("shopping_cart_id", "shopping_carts(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.DeliveryPayment{}).AddForeignKey("delivery_id", "deliveries(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.DeliveryPayment{}).AddForeignKey("payment_id", "payments(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Favorite{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Feedback{}).AddForeignKey("product_id", "products(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Media{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.PaymentInfo{}).AddForeignKey("payment_id", "payments(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Permission{}).AddForeignKey("auth_id", "auths(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("category_id", "categories(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.ProductTranslation{}).AddForeignKey("product_id", "products(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserPlan{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserPlan{}).AddForeignKey("plan_id", "plans(uuid)", "RESTRICT", "RESTRICT")
}
