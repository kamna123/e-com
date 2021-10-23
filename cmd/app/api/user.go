package api

import (
	"e-commerce/cmd/app/schema"
	"e-commerce/cmd/app/services"
	"e-commerce/cmd/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type User struct {
	service services.IUserService
}

func NewUserAPI(service services.IUserService) *User {
	return &User{service: service}
}

func (u *User) validate(r schema.Register) bool {
	return utils.Validate(
		[]utils.Validation{
			{Value: r.Username, Valid: "username"},
			{Value: r.Email, Valid: "email"},
			{Value: r.Password, Valid: "password"},
		})
}

// Login godoc
// @Summary Login user
// @Produce json
// @Accept json
// @Param Body body schema.Login true "The body to create a login details"
// @Security ApiKeyAuth
// @Success 200 {object} schema.User
// @Router /auth/auth/login [post]
func (u *User) Login(c *gin.Context) {
	var item schema.Login
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	user, token, err := u.service.Login(ctx, &item)
	if err != nil {
		///log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.User
	copier.Copy(&res, &user)
	res.Extra = map[string]interface{}{
		"token": token,
	}
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

// Login godoc
// @Summary Register user
// @Produce json
// @Accept json
// @Param Body body schema.Register true "The body to register a user"
// @Security ApiKeyAuth
// @Success 200 {object} schema.User
// @Router /auth/auth/register  [post]
func (u *User) Register(c *gin.Context) {
	var item schema.Register
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := u.validate(item)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is invalid"})
		return
	}

	ctx := c.Request.Context()
	user, token, err := u.service.Register(ctx, &item)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.User
	copier.Copy(&res, &user)
	res.Extra = map[string]interface{}{
		"token": token,
	}
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
