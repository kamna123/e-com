package repositories

import (
	"errors"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/database"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	AddToCart(query *schema.CartBody) (*models.Cart, error)
	GetCart(uuid string) (*[]models.Cart, error)
	UpdateFromCart(query *schema.CartDeleteBody) (*models.Cart, error)
}

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepository() ICartRepository {
	return &cartRepo{db: database.Database}
}

func (r *cartRepo) AddToCart(item *schema.CartBody) (*models.Cart, error) {
	var cart models.Cart
	copier.Copy(&cart, &item)

	if err := r.db.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepo) GetCart(uuid string) (*[]models.Cart, error) {
	var cart []models.Cart
	if r.db.Where("user_id = ?", uuid).Find(&cart).RecordNotFound() {
		return nil, errors.New("not found cart")
	}

	return &cart, nil
}

func (r *cartRepo) UpdateFromCart(query *schema.CartDeleteBody) (*models.Cart, error) {
	var product models.Cart
	if r.db.Where("user_id = ? and product_uuid = ?", query.UserID, query.ProductUUID).Find(&product).RecordNotFound() {
		return nil, errors.New("not found cart")
	}
	if r.db.Model(&product).Where("user_id = ? and product_uuid = ? ", query.UserID, query.ProductUUID).
		Update("quantity", query.Quantity).RecordNotFound() {
		return nil, errors.New("not found product")
	}

	return &product, nil
}
