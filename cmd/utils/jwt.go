package utils

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/copier"
)

const (
	TokenExpiredTime = 300
)

func GenerateToken(payload interface{}) string {
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * TokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	if err != nil {
		glog.Error("Failed to generate token: ", err)
		return ""
	}

	return token
}

func ValidateToken(jwtToken string) (map[string]interface{}, error) {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	var data map[string]interface{}
	copier.Copy(&data, tokenData["payload"])
	return data, nil
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code string

		code = Success
		token := c.GetHeader("Authorization")

		if token == "" {
			code = InvalidParams
			c.JSON(http.StatusUnauthorized, PrepareResponse(nil, "Unauthorized", code))

			c.Abort()
			return
		}

		_, err := ValidateToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = ErrorAuthCheckTokenTimeout
			default:
				code = ErrorAuthCheckTokenFail
			}
		}

		if code != Success {
			c.JSON(http.StatusUnauthorized, PrepareResponse(nil, "Unauthorized", code))

			c.Abort()
			return
		}

		c.Next()
	}
}
