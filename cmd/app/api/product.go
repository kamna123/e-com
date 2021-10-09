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

type Product struct {
	service services.IProductService
}

func NewProductAPI(service services.IProductService) *Product {
	return &Product{service: service}
}

// GetProductByID godoc
// @Summary Get get product by uuid
// @Produce json
// @Param uuid path string true "Product UUID"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Product
// @Router /api/v1/products/{uuid} [get]
func (p *Product) GetProductByID(c *gin.Context) {
	productId := c.Param("uuid")

	ctx := c.Request.Context()
	product, err := p.service.GetProductByID(ctx, productId)
	if err != nil {
		glog.Error("Failed to get product: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Product
	copier.Copy(&res, &product)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// GetProducts godoc
// @Summary Get list products
// @Produce json
// @Param Body body schema.ProductQueryParam true "The body to get orders"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.Product
// @Router /api/v1/products [get]
func (categ *Product) GetProducts(c *gin.Context) {
	var params schema.ProductQueryParam
	if err := c.ShouldBindQuery(&params); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := categ.service.GetProducts(ctx, params)
	if err != nil {
		glog.Error("Failed to get categories: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Product
	copier.Copy(&res, &rs)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// GetProductByCategoryID godoc
// @Summary Get get product by category ID
// @Produce json
// @Param uuid path string true "Product UUID"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.Product
// @Router /api/v1/products/{uuid} [get]
func (p *Product) GetProductByCategoryID(c *gin.Context) {
	categUUID := c.Param("uuid")

	ctx := c.Request.Context()
	products, err := p.service.GetProductByCategoryID(ctx, categUUID)
	if err != nil {
		glog.Error("Failed to get products: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Product
	copier.Copy(&res, &products)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// CreateProduct godoc
// @Summary Post create product
// @Produce json
// @Accept json
// @Param Body body schema.ProductBodyParam true "The body to create a product"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.Product
// @Router /api/v1/products [post]
func (p *Product) CreateProduct(c *gin.Context) {
	var item schema.ProductBodyParam
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
	products, err := p.service.CreateProduct(ctx, &item)
	if err != nil {
		glog.Error("Failed to create product", err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Product
	copier.Copy(&res, &products)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// UpdateProduct godoc
// @Summary Put update product
// @Produce json
// @Accept json
// @Param Body body schema.ProductBodyParam true "The body to update a product"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Product
// @Router /api/v1/products [put]
func (p *Product) UpdateProduct(c *gin.Context) {
	uuid := c.Param("uuid")
	var item schema.ProductBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		glog.Error("Failed to parse request body: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	products, err := p.service.UpdateProduct(ctx, uuid, &item)
	if err != nil {
		glog.Error("Failed to update product: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Product
	copier.Copy(&res, &products)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
