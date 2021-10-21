package services

import (
	"context"
	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/repositories"
	"e-commerce/cmd/app/schema"

	"github.com/golang/glog"
)

type ICartService interface {
	AddToCart(ctx context.Context, query *schema.CartBody) (*models.Cart, error)
	GetCart(ctx context.Context, uuid string) (*[]models.Cart, error)
	UpdateFromCart(ctx context.Context, query *schema.CartDeleteBody) (*models.Cart, error)
}

type cart struct {
	repo repositories.ICartRepository
}

func NewCartService(repo repositories.ICartRepository) ICartService {
	return &cart{repo: repo}
}

func (c *cart) AddToCart(ctx context.Context, query *schema.CartBody) (*models.Cart, error) {
	categories, err := c.repo.AddToCart(query)
	if err != nil {
		glog.Error("Failed to get categories: ", err)
		return nil, err
	}

	return categories, nil
}

func (c *cart) GetCart(ctx context.Context, uuid string) (*[]models.Cart, error) {
	category, err := c.repo.GetCart(uuid)
	if err != nil {
		glog.Error("Failed to get cart: ", err)
		return nil, err
	}

	return category, nil
}

func (c *cart) UpdateFromCart(ctx context.Context, query *schema.CartDeleteBody) (*models.Cart, error) {
	category, err := c.repo.UpdateFromCart(query)
	if err != nil {
		glog.Error("Failed to get cart: ", err)
		return nil, err
	}

	return category, nil
}
