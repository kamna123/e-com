package repositories

import (
	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/database"
	"e-commerce/cmd/utils"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Login(item *schema.Login) (*models.User, error)
	Register(item *schema.Register) (*models.User, error)
	// GetUserByID(uuid string) (*models.User, error)
}
type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) Login(item *schema.Login) (*models.User, error) {
	user := &models.User{}
	if database.Database.Where("username = ? ", item.Username).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(item.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
func (u *userRepo) Register(item *schema.Register) (*models.User, error) {
	var user models.User
	copier.Copy(&user, &item)
	hashedPassword := utils.HashAndSalt([]byte(item.Password))
	user.Password = hashedPassword

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository() UserRepository {
	return &userRepo{db: database.Database}
}
