package router

import (
	"e-commerce/cmd/app/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *api.User,
		role *api.Role,
	) {

	})
	return err
}
