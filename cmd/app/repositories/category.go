package repositories

import (
	"errors"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/database"
	"e-commerce/cmd/utils"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	GetCategories(query *schema.CategoryQueryParam) (*[]models.Category, error)
	GetCategoryByID(uuid string) (*models.Category, error)
	CreateCategory(item *schema.Category) (*models.Category, error)
}

type categRepo struct {
	db *gorm.DB
}

func NewCategoryRepository() ICategoryRepository {
	return &categRepo{db: database.Database}
}

func (r *categRepo) GetCategories(query *schema.CategoryQueryParam) (*[]models.Category, error) {
	var categories []models.Category
	var mapQuery map[string]interface{}
	utils.Copy(&mapQuery, query)
	if r.db.Find(&categories).RecordNotFound() {
		return nil, nil
	}

	return &categories, nil
}

func (r *categRepo) GetCategoryByID(uuid string) (*models.Category, error) {
	var category models.Category
	if r.db.Where("uuid = ?", uuid).Find(&category).RecordNotFound() {
		return nil, errors.New("not found category")
	}

	return &category, nil
}

func (r *categRepo) CreateCategory(item *schema.Category) (*models.Category, error) {
	var category models.Category
	copier.Copy(&category, &item)

	if err := r.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
