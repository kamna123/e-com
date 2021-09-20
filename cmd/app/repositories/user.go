package repositories

import (
	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Login(item *schema.Login) (*models.User, error)
	Register(item *schema.Register) (*models.User, error)
	GetUserByID(uuid string) (*models.User, error)
}
type userRepo struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{db: dbs.Database}
}
