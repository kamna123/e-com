package models

import (
	"e-commerce/cmd/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	UUID          string `json:"uuid" gorm:"unique;not null;index;primaryKey"`
	Code          string `json:"code" gorm:"unique;not null;index"`
	Name          string `json:"name"`
	Brand         string `json:"brand"`
	Description   string `json:"description"`
	CategUUID     string `json:"categ_uuid"`
	Price         uint   `json:"price"`
	Active        bool   `json:"active" gorm:"default:true"`
	VideoPath     string `json:"video_path"`
	ThumbnailPath string `json:"thumbnail_path"`
	gorm.Model
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	product.UUID = uuid.New().String()
	product.Code = utils.GenerateCode("P")
	product.Active = true
	return nil
}
