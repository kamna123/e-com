// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/cart": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post add item to cart",
                "parameters": [
                    {
                        "description": "The body to create a order",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CartBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.CartBody"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/cart/delete": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Put delete item from cart",
                "parameters": [
                    {
                        "description": "The body to update a product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CartDeleteBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CartBody"
                        }
                    }
                }
            }
        },
        "/api/v1/cart/{uuid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get cart by user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart user id",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CartBody"
                        }
                    }
                }
            }
        },
        "/api/v1/categories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get category by code",
                "parameters": [
                    {
                        "description": "The body to get categories",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CategoryQueryParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Category"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/categories/{uuid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get category by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Category"
                        }
                    }
                }
            }
        },
        "/api/v1/orders": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get order by query param",
                "parameters": [
                    {
                        "description": "The body to get orders",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OrderQueryParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post create order",
                "parameters": [
                    {
                        "description": "The body to create a order",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OrderBodyParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Order"
                        }
                    }
                }
            }
        },
        "/api/v1/orders/{uuid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get order by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Order"
                        }
                    }
                }
            }
        },
        "/api/v1/products": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Product"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Put update product",
                "parameters": [
                    {
                        "description": "The body to update a product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.ProductBodyParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Product"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post create product",
                "parameters": [
                    {
                        "description": "The body to create a product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.ProductBodyParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Product"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/products/{uuid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get product by category ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Product"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/quantities": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get quantities",
                "parameters": [
                    {
                        "description": "The body to get categories",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.QuantityQueryParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Quantity"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Put update quantity",
                "parameters": [
                    {
                        "description": "The body to create a order",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.QuantityBodyParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Quantity"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post create quantity",
                "parameters": [
                    {
                        "description": "The body to create a order",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.QuantityBodyParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Quantity"
                        }
                    }
                }
            }
        },
        "/api/v1/quantities/{uuid}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get get quantity by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quantity UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Quantity"
                        }
                    }
                }
            }
        },
        "/auth/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "The body to create a login details",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    }
                }
            }
        },
        "/auth/auth/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "The body to register a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.CartBody": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "string"
                },
                "userid": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schema.CartDeleteBody": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "string"
                },
                "userid": {
                    "type": "string"
                }
            }
        },
        "schema.Category": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schema.CategoryQueryParam": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "code": {
                    "type": "string"
                }
            }
        },
        "schema.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.Order": {
            "type": "object",
            "properties": {
                "lines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.OrderLine"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total_price": {
                    "type": "integer"
                },
                "userid": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schema.OrderBodyParam": {
            "type": "object",
            "required": [
                "lines",
                "userid"
            ],
            "properties": {
                "lines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.OrderLineBodyParam"
                    }
                },
                "userid": {
                    "type": "string"
                }
            }
        },
        "schema.OrderLine": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "integer"
                },
                "product_uuid": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schema.OrderLineBodyParam": {
            "type": "object",
            "required": [
                "product_uuid",
                "quantity"
            ],
            "properties": {
                "product_uuid": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "schema.OrderQueryParam": {
            "type": "object",
            "required": [
                "userid"
            ],
            "properties": {
                "status": {
                    "type": "string"
                },
                "userid": {
                    "type": "string"
                }
            }
        },
        "schema.Product": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "categ_uuid": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "schema.ProductBodyParam": {
            "type": "object",
            "required": [
                "categ_uuid",
                "name"
            ],
            "properties": {
                "categ_uuid": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "schema.Quantity": {
            "type": "object",
            "properties": {
                "product_uuid": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                },
                "warehouse_uuid": {
                    "type": "string"
                }
            }
        },
        "schema.QuantityBodyParam": {
            "type": "object",
            "required": [
                "product_uuid",
                "quantity",
                "warehouse_uuid"
            ],
            "properties": {
                "product_uuid": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "warehouse_uuid": {
                    "type": "string"
                }
            }
        },
        "schema.QuantityQueryParam": {
            "type": "object",
            "properties": {
                "product_uuid": {
                    "type": "string"
                },
                "warehouse_uuid": {
                    "type": "string"
                }
            }
        },
        "schema.Register": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role_uuid": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "extra": {},
                "username": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}