package migrations

import (
	"e-commerce/cmd/app/models"
	dbs "e-commerce/cmd/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Migrate() {
	User := models.User{}
	Category := models.Category{}
	Product := models.Product{}
	Order := models.Order{}
	OrderLine := models.OrderLine{}
	Cart := models.Cart{}
	Quantity := models.Quantity{}
	dbs.Database.LogMode(true)
	//dbs.Database.AutoMigrate(&Product)
	dbs.Database.AutoMigrate(&User, &Category, &Product, &Order, &OrderLine, &Cart, &Quantity)
	dbs.Database.Model(&Product).AddForeignKey("categ_uuid", "categories(uuid)", "RESTRICT", "RESTRICT")
	dbs.Database.Model(&Order).AddForeignKey("user_id", "users(uuid)", "RESTRICT", "RESTRICT")
	dbs.Database.Model(&OrderLine).AddForeignKey("product_uuid", "products(uuid)", "RESTRICT", "RESTRICT")
	dbs.Database.Model(&OrderLine).AddForeignKey("order_uuid", "orders(uuid)", "RESTRICT", "RESTRICT")
	dbs.Database.Model(&Cart).AddForeignKey("user_id", "users(uuid)", "RESTRICT", "RESTRICT")
	dbs.Database.Model(&Cart).AddForeignKey("product_uuid", "products(uuid)", "RESTRICT", "RESTRICT")
	dbs.Database.Model(&Quantity).AddForeignKey("product_uuid", "products(uuid)", "RESTRICT", "RESTRICT")

	//createAdmin()
}
