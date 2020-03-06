package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
)

func MigrateAssociation() {
	db := helper.GetDB()
	// db.Model(&models.Auth{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Category{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.CategoryTranslation{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Contact{}).AddForeignKey("id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Delivery{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.DeliveryPayment{}).AddForeignKey("delivery_id", "deliveries(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.DeliveryPayment{}).AddForeignKey("payment_id", "payments(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Favorite{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Feedback{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Media{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.PaymentInfo{}).AddForeignKey("payment_id", "payments(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Permission{}).AddForeignKey("auth_id", "auths(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.ProductTranslation{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserPlan{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserPlan{}).AddForeignKey("plan_id", "plans(id)", "RESTRICT", "RESTRICT")
}
