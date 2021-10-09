package api

import (
	"net/http"

	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/app/services"
	"e-commerce/cmd/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/glog"
	"github.com/jinzhu/copier"
)

type Order struct {
	service services.IOrderSerivce
}

func NewOrderAPI(service services.IOrderSerivce) *Order {
	return &Order{service: service}
}

// CreateOrder godoc
// @Summary Post get razor pay order id
// @Produce json
// @Accept json
// @Param Body body schema.RazorPayOrderParam true "The body to create a order"
// @Security ApiKeyAuth
// @Success 200 {object} schema.RazorPayResp
// @Router /api/v1/order/razorpay [post]
func (categ *Order) RazorPayOrder(c *gin.Context) {
	var query schema.RazorPayOrderParam
	if err := c.Bind(&query); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	orders, err := categ.service.RazorPayOrder(ctx, &query)
	if err != nil {
		glog.Error("Failed to get orders: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.RazorPayResp
	copier.Copy(&res, &orders)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))

}

// GetOrders godoc
// @Summary Get get order by query param
// @Produce json
// @Accept json
// @Param Body body schema.OrderQueryParam true "The body to get orders"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.Order
// @Router /api/v1/orders [get]
func (categ *Order) GetOrders(c *gin.Context) {
	var query schema.OrderQueryParam
	if err := c.ShouldBindQuery(&query); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	orders, err := categ.service.GetOrders(ctx, &query)
	if err != nil {
		glog.Error("Failed to get orders: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Order
	copier.Copy(&res, &orders)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// GetOrderByID godoc
// @Summary Get get order by uuid
// @Produce json
// @Param uuid path string true "Order UUID"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Order
// @Router /api/v1/orders/{uuid} [get]
func (categ *Order) GetOrderByID(c *gin.Context) {
	orderId := c.Param("uuid")

	ctx := c.Request.Context()
	order, err := categ.service.GetOrderByID(ctx, orderId)
	if err != nil {
		glog.Error("Failed to get Order: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Order
	copier.Copy(&res, &order)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// CreateOrder godoc
// @Summary Post create order
// @Produce json
// @Accept json
// @Param Body body schema.OrderBodyParam true "The body to create a order"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Order
// @Router /api/v1/orders [post]
func (categ *Order) CreateOrder(c *gin.Context) {
	var item schema.OrderBodyParam
	if err := c.Bind(&item); err != nil {
		glog.Error("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(item)
	if err != nil {
		glog.Error("Request body is invalid: ", err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	ctx := c.Request.Context()
	orders, err := categ.service.CreateOrder(ctx, &item)
	if err != nil {
		glog.Error("Failed to create Order: ", err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Order
	copier.Copy(&res, &orders)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
