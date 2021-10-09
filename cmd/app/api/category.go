package api

import (
	"net/http"

	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/app/services"
	"e-commerce/cmd/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/copier"
)

type Category struct {
	service services.ICategoryService
}

func NewCategoryAPI(service services.ICategoryService) *Category {
	return &Category{service: service}
}

// GetCategoryByID godoc
// @Summary Get get category by code
// @Produce json
// @Accept json
// @Param Body body schema.CategoryQueryParam true "The body to get categories"
// @Security ApiKeyAuth
// @Success 200 {object} []schema.Category
// @Router /api/v1/categories [get]
func (categ *Category) GetCategories(c *gin.Context) {
	var query schema.CategoryQueryParam
	if err := c.ShouldBindQuery(&query); err != nil {
		glog.Error("Failed to parse request query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := categ.service.GetCategories(ctx, &query)
	if err != nil {
		glog.Error("Failed to get categories: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res []schema.Category
	copier.Copy(&res, &rs)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// GetCategoryByID godoc
// @Summary Get get category by uuid
// @Produce json
// @Param uuid path string true "Category UUID"
// @Security ApiKeyAuth
// @Success 200 {object} schema.Category
// @Router /api/v1/categories/{uuid} [get]
func (categ *Category) GetCategoryByID(c *gin.Context) {
	categoryId := c.Param("uuid")

	ctx := c.Request.Context()
	category, err := categ.service.GetCategoryByID(ctx, categoryId)
	if err != nil {
		glog.Error("Failed to get category: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Category
	copier.Copy(&res, &category)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
