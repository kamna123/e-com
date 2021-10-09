package repositories

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"
	dbs "e-commerce/cmd/database"
)

type AddressRepository interface {
	GetAddressByUserID(uuid string) (*models.Address, error)
	GetAddressByID(uuid string) (*models.Address, error)

	CreateAddress(item *schema.Address) (*models.Address, error)
	UpdateAddress(uuid string, item *schema.Address) (*models.Address, error)
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepository() AddressRepository {
	return &addressRepo{db: dbs.Database}
}

func (w *addressRepo) GetAddressByUserID(uuid string) (*models.Address, error) {
	var warehouse models.Address
	if w.db.Where("user_id = ?", uuid).First(&warehouse).RecordNotFound() {
		return nil, errors.New("not found address")
	}

	return &warehouse, nil
}

func (w *addressRepo) CreateAddress(item *schema.Address) (*models.Address, error) {
	var warehouse models.Address
	copier.Copy(&warehouse, &item)

	if err := w.db.Create(&warehouse).Error; err != nil {
		return nil, err
	}

	return &warehouse, nil
}
func (w *addressRepo) GetAddressByID(uuid string) (*models.Address, error) {
	var warehouse models.Address
	if w.db.Where("uuid = ?", uuid).First(&warehouse).RecordNotFound() {
		return nil, errors.New("not found warehouse")
	}

	return &warehouse, nil
}

func (w *addressRepo) UpdateAddress(uuid string, item *schema.Address) (*models.Address, error) {
	warehouse, err := w.GetAddressByID(uuid)
	if err != nil {
		return nil, err
	}

	copier.Copy(warehouse, &item)
	if err := w.db.Save(&warehouse).Error; err != nil {
		return nil, err
	}

	return warehouse, nil
}
