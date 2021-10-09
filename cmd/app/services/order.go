package services

import (
	"context"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"

	"e-commerce/cmd/app/repositories"
)

type IOrderSerivce interface {
	GetOrders(ctx context.Context, query *schema.OrderQueryParam) (*[]models.Order, error)
	GetOrderByID(ctx context.Context, uuid string) (*models.Order, error)
	CreateOrder(ctx context.Context, item *schema.OrderBodyParam) (*models.Order, error)
}

type order struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) IOrderSerivce {
	return &order{repo: repo}
}

func (categ *order) GetOrders(ctx context.Context, query *schema.OrderQueryParam) (*[]models.Order, error) {
	orders, err := categ.repo.GetOrders(query)
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (categ *order) GetOrderByID(ctx context.Context, uuid string) (*models.Order, error) {
	order, err := categ.repo.GetOrderByID(uuid)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (categ *order) CreateOrder(ctx context.Context, item *schema.OrderBodyParam) (*models.Order, error) {
	order, err := categ.repo.CreateOrder(item)
	if err != nil {
		return nil, err
	}

	return order, nil
}
