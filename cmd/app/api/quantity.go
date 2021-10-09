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

type Quantity struct {
	service services.IQuantityService
}

func NewQuantityAPI(service services.IQuantityService) *Quantity {
	return &Quantity{service: service}
}

// GetQuantities godoc
// @Summary Get get quantities
// @Produce json
// @Accept json
// @Param Body body schema.QuantityQueryParam true "The body to get categories"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.Quantity
// @Router /api/v1/quantities [get]
func (q *Quantity) GetQuantities(c *gin.Context) {
	var query schema.QuantityQueryParam
	if err := c.ShouldBindQuery(&query); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	quantities, err := q.service.GetQuantities(ctx, &query)
	if err != nil {
		glog.Error("Failed to get quantities}: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Quantity
	copier.Copy(&res, &quantities)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// GetQuantityByID godoc
// @Summary Get get quantity by uuid
// @Produce json
// @Param uuid path string true "Quantity UUID"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Quantity
// @Router /api/v1/quantities/{uuid} [get]
func (q *Quantity) GetQuantityByID(c *gin.Context) {
	quantityId := c.Param("uuid")

	ctx := c.Request.Context()
	quantity, err := q.service.GetQuantityByID(ctx, quantityId)
	if err != nil {
		glog.Error("Failed to get quantity: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Quantity
	copier.Copy(&res, &quantity)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// CreateQuantity godoc
// @Summary Post create quantity
// @Produce json
// @Accept json
// @Param Body body schema.QuantityBodyParam true "The body to create a order"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Quantity
// @Router /api/v1/quantities  [post]
func (q *Quantity) CreateQuantity(c *gin.Context) {
	var item schema.QuantityBodyParam
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
	quantities, err := q.service.CreateQuantity(ctx, &item)
	if err != nil {
		glog.Error("Failed to create quantity", err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Quantity
	copier.Copy(&res, &quantities)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// UpdateQuantity godoc
// @Summary Put update quantity
// @Produce json
// @Accept json
// @Param Body body schema.QuantityBodyParam true "The body to create a order"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Quantity
// @Router /api/v1/quantities  [put]
func (q *Quantity) UpdateQuantity(c *gin.Context) {
	uuid := c.Param("uuid")
	var item schema.QuantityBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		glog.Error("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	quantities, err := q.service.UpdateQuantity(ctx, uuid, &item)
	if err != nil {
		glog.Error("Failed to update quantity: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Quantity
	copier.Copy(&res, &quantities)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
