package router

import (
	"e-commerce/cmd/app/api"

	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *api.User,
		category *api.Category,
		//role *api.Role,
	) error {
		auth := r.Group("/auth")
		{
			auth.POST("auth/register", user.Register)
			auth.POST("auth/login", user.Login)

		}
		apiV1 := r.Group("api/v1")
		{
			apiV1.GET("/categories", category.GetCategories)
			apiV1.GET("/categories/:uuid", category.GetCategoryByID)

		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return err
}
