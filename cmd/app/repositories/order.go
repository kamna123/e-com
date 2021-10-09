package repositories

import (
	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"
	dbs "e-commerce/cmd/database"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	GetOrders(query *schema.OrderQueryParam) (*[]models.Order, error)
	GetOrderByID(uuid string) (*models.Order, error)
	CreateOrder(item *schema.OrderBodyParam) (*models.Order, error)
	RazorPayOrder() (string, error)
}

type orderRepo struct {
	db           *gorm.DB
	lineRepo     OrderLineRepository
	quantityRepo QuantityRepository
}

func NewOrderRepository() OrderRepository {
	return &orderRepo{db: dbs.Database, lineRepo: NewOrderLineRepository()}
}

func (r *orderRepo) RazorPayOrder() (string, error) {
	return "", nil
}
func (r *orderRepo) GetOrders(query *schema.OrderQueryParam) (*[]models.Order, error) {
	var orders []models.Order
	if r.db.Find(&orders, query).RecordNotFound() {
		return nil, nil
	}

	return &orders, nil
}

func (r *orderRepo) GetOrderByID(uuid string) (*models.Order, error) {
	var order models.Order
	var lines []models.OrderLine
	if r.db.Where("uuid = ?", uuid).First(&order).RecordNotFound() {
		return nil, errors.New("not found order")
	}
	r.db.Where("order_uuid = ?", uuid).Find(&lines)
	order.Lines = lines

	return &order, nil
}

func (r *orderRepo) CreateOrder(item *schema.OrderBodyParam) (*models.Order, error) {
	var order models.Order
	copier.Copy(&order, &item)

	if order.Lines == nil {
		return nil, errors.New("order lines must be not empty")
	}

	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}

	var lines []models.OrderLine
	var totalPrice uint
	for _, line := range order.Lines {
		line.OrderUUID = order.UUID
		if err := r.db.Create(&line).Error; err != nil {
			return nil, err
		}
		lines = append(lines, line)
		totalPrice += line.Price
	}
	order.TotalPrice = totalPrice
	order.Lines = lines

	if err := r.db.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
