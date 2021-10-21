package api

import (
	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/app/services"
	"e-commerce/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/copier"
)

type Cart struct {
	service services.ICartService
}

func NewCartAPI(service services.ICartService) *Cart {
	return &Cart{service: service}
}

// AddToCart godoc
// @Summary Post add item to cart
// @Produce json
// @Accept json
// @Param Body body schema.CartBody true "The body to create a order"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.CartBody
// @Router /api/v1/cart [post]
func (categ *Cart) AddToCart(c *gin.Context) {
	var query schema.CartBody
	if err := c.Bind(&query); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := categ.service.AddToCart(ctx, &query)
	if err != nil {
		glog.Error("Failed to get categories: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.CartBody
	copier.Copy(&res, &rs)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// GetCart godoc
// @Summary Get get cart by user id
// @Produce json
// @Param uuid path string true "cart user id"
// @Security ApiKeyAuth
// @Success 200 {object} schema.CartBody
// @Router /api/v1/cart/{uuid} [get]
func (categ *Cart) GetCart(c *gin.Context) {
	categoryId := c.Param("uuid")

	ctx := c.Request.Context()
	category, err := categ.service.GetCart(ctx, categoryId)
	if err != nil {
		glog.Error("Failed to get category: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.CartBody
	copier.Copy(&res, &category)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// DeleteFromCart godoc
// @Summary Put delete item from cart
// @Produce json
// @Accept json
// @Param Body body schema.CartDeleteBody true "The body to update a product"
// @Security ApiKeyAuth
// @Success 200 {object} schema.CartBody
// @Router /api/v1/cart/update [put]
func (categ *Cart) UpdateFromCart(c *gin.Context) {
	var query schema.CartDeleteBody
	if err := c.Bind(&query); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	category, err := categ.service.UpdateFromCart(ctx, &query)
	if err != nil {
		glog.Error("Failed to get category: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.CartBody
	copier.Copy(&res, &category)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
