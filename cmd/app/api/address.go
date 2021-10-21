package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/glog"
	"github.com/jinzhu/copier"

	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/app/services"
	"e-commerce/cmd/utils"
)

type Address struct {
	service services.IAddressSerivce
}

func NewAddressAPI(service services.IAddressSerivce) *Address {
	return &Address{service: service}
}

// GetAddressByUserID godoc
// @Summary Get get address by userid
// @Produce json
// @Param uuid path string true "user id"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Address
// @Router /api/v1/address/{uuid} [get]
func (w *Address) GetAddressByUserID(c *gin.Context) {
	AddressId := c.Param("uuid")

	ctx := c.Request.Context()
	Address, err := w.service.GetAddressByUserID(ctx, AddressId)
	if err != nil {
		glog.Error("Failed to get Address: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Address
	copier.Copy(&res, &Address)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// CreateAddress godoc
// @Summary Post create address
// @Produce json
// @Accept json

// @Param Body body schema.Address true "The body to create a order"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Address
// @Router /api/v1/address [post]
func (w *Address) CreateAddress(c *gin.Context) {
	var item schema.Address
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
	Addresss, err := w.service.CreateAddress(ctx, &item)
	if err != nil {
		glog.Error("Failed to create Address", err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Address
	copier.Copy(&res, &Addresss)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// UpdateAddress godoc
// @Summary Put update address
// @Produce json
// @Accept json
// @Param uuid path string true "user id"
// @Param Body body schema.Address true "The body to update a address"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Address
// @Router /api/v1/address/{uuid} [put]
func (w *Address) UpdateAddress(c *gin.Context) {
	uuid := c.Param("uuid")
	var item schema.Address
	if err := c.ShouldBindJSON(&item); err != nil {
		glog.Error("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	Addresss, err := w.service.UpdateAddress(ctx, uuid, &item)
	if err != nil {
		glog.Error("Failed to update Address: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Address
	copier.Copy(&res, &Addresss)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
