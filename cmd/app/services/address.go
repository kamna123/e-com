package services

import (
	"context"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/repositories"
	"e-commerce/cmd/app/schema"
)

type IAddressSerivce interface {
	GetAddressByUserID(ctx context.Context, uuid string) (*models.Address, error)
	CreateAddress(ctx context.Context, item *schema.Address) (*models.Address, error)
	UpdateAddress(ctx context.Context, uuid string, items *schema.Address) (*models.Address, error)
}

type addressRepo struct {
	repo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) IAddressSerivce {
	return &addressRepo{repo: repo}
}

func (w *addressRepo) GetAddressByUserID(ctx context.Context, uuid string) (*models.Address, error) {
	warehouse, err := w.repo.GetAddressByUserID(uuid)
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (w *addressRepo) CreateAddress(ctx context.Context, item *schema.Address) (*models.Address, error) {
	warehouse, err := w.repo.CreateAddress(item)
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (w *addressRepo) UpdateAddress(ctx context.Context, uuid string, item *schema.Address) (*models.Address, error) {
	warehouse, err := w.repo.UpdateAddress(uuid, item)
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}
