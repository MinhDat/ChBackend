package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func MigrateAssociation() {
	db := helper.GetDB()
	db.Model(&models.Auth{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Category{}).AddForeignKey("owner_id", "users(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Category{}).AddForeignKey("parent_id", "categories(uuid)", "SET NULL", "CASCADE")
	db.Model(&models.CategoryTranslation{}).AddForeignKey("category_id", "categories(uuid)", "RESTRICT", "RESTRICT")
	db.Model(&models.Contact{}).AddForeignKey("uuid", "users(uuid)", "CASCADE", "CASCADE")
	db.Model(&models.Delivery{}).AddForeignKey("order_id", "orders(uuid)", "RESTRICT", "RESTRICT")
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
