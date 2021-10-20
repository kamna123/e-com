package services

import (
	"context"
	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/repositories"
	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/utils"
)

type IUserService interface {
	Login(ctx context.Context, item *schema.Login) (*models.User, string, error)
	Register(ctx context.Context, item *schema.Register) (*models.User, string, error)
	// GetUserByID(ctx context.Context, uuid string) (*models.User, error)
}
type user struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) IUserService {
	return &user{repo: repo}
}

func (u *user) Login(ctx context.Context, item *schema.Login) (*models.User, string, error) {
	user, err := u.repo.Login(item)
	if err != nil {
		return nil, "", err
	}

	token := utils.GenerateToken(user)
	return user, token, nil
}

func (u *user) Register(ctx context.Context, item *schema.Register) (*models.User, string, error) {
	user, err := u.repo.Register(item)
	if err != nil {
		return nil, "", err
	}

	token := utils.GenerateToken(user)
	return user, token, nil
}
