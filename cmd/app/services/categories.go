package services

import (
	"context"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"

	"github.com/golang/glog"

	"e-commerce/cmd/app/repositories"
)

type ICategoryService interface {
	GetCategories(ctx context.Context, query *schema.CategoryQueryParam) (*[]models.Category, error)
	GetCategoryByID(ctx context.Context, uuid string) (*models.Category, error)
}

type category struct {
	repo repositories.ICategoryRepository
}

func NewCategoryService(repo repositories.ICategoryRepository) ICategoryService {
	return &category{repo: repo}
}

func (c *category) GetCategories(ctx context.Context, query *schema.CategoryQueryParam) (*[]models.Category, error) {
	categories, err := c.repo.GetCategories(query)
	if err != nil {
		glog.Error("Failed to get categories: ", err)
		return nil, err
	}

	return categories, nil
}

func (c *category) GetCategoryByID(ctx context.Context, uuid string) (*models.Category, error) {
	category, err := c.repo.GetCategoryByID(uuid)
	if err != nil {
		glog.Error("Failed to get category: ", err)
		return nil, err
	}

	return category, nil
}
