package api

import (
	"e-commerce/cmd/app/services"
)

type User struct {
	service services.IUserService
}
