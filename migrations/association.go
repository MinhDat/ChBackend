package migrate

import (
	helper "ChGo/helpers"
	"ChGo/models"
)

func MigrateAssociation() {
	db := helper.GetDB()
	db.Model(&models.Category{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.CategoryTranslation{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&models.Contact{}).AddForeignKey("id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Delivery{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	db.Model(&models.DeliveryPayment{}).AddForeignKey("delivery_id", "deliveries(id)", "CASCADE", "CASCADE")
	db.Model(&models.DeliveryPayment{}).AddForeignKey("payment_id", "payments(id)", "CASCADE", "NO ACTION")
	db.Model(&models.Favorite{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Feedback{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&models.Media{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.PaymentInfo{}).AddForeignKey("payment_id", "payments(id)", "CASCADE", "CASCADE")
	db.Model(&models.Permission{}).AddForeignKey("auth_id", "auths(id)", "CASCADE", "CASCADE")
	db.Model(&models.Product{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Product{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&models.ProductTranslation{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&models.UserAuth{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.UserAuth{}).AddForeignKey("auth_id", "auths(name)", "CASCADE", "NO ACTION")
	db.Model(&models.UserPlan{}).AddForeignKey("owner_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.UserPlan{}).AddForeignKey("plan_id", "plans(id)", "CASCADE", "NO ACTION")
}
