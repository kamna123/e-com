package services

import (
	"context"
	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"
)

type IUserService interface {
	Login(ctx context.Context, item *schema.Login) (*models.User, string, error)
	Register(ctx context.Context, item *schema.Register) (*models.User, string, error)
	GetUserByID(ctx context.Context, uuid string) (*models.User, error)
}
