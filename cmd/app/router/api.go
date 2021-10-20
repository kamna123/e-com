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
		product *api.Product,
		quantity *api.Quantity,
		cart *api.Cart,
		order *api.Order,
		address *api.Address,
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
		{
			apiV1.GET("/products", product.GetProducts)
			apiV1.POST("/products", product.CreateProduct)
			apiV1.GET("/products/:uuid", product.GetProductByCategoryID)
			apiV1.PUT("/products/:uuid", product.UpdateProduct)
		}
		{
			apiV1.GET("/quantities", quantity.GetQuantities)
			apiV1.POST("/quantities", quantity.CreateQuantity)
			apiV1.GET("/quantities/:uuid", quantity.GetQuantityByID)
			apiV1.PUT("/quantities/:uuid", quantity.UpdateQuantity)
		}
		{
			apiV1.GET("/cart/:uuid", cart.GetCart)
			apiV1.POST("/cart", cart.AddToCart)
			apiV1.PUT("/cart/delete", cart.DeleteFromCart)
		}
		{
			apiV1.GET("/orders", order.GetOrders)
			apiV1.POST("/orders", order.CreateOrder)
			apiV1.POST("/order/razorpay", order.RazorPayOrder)
			apiV1.GET("/orders/:uuid", order.GetOrderByID)
		}
		{
			apiV1.GET("/address/:uuid", address.GetAddressByUserID)
			apiV1.POST("/address", address.CreateAddress)
			apiV1.PUT("/address/:uuid", address.UpdateAddress)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return err
}
