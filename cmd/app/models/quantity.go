package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Quantity struct {
	UUID        string `json:"uuid" gorm:"unique;not null;index;primaryKey"`
	ProductUUID string `json:"product_uuid" gorm:"not null;index"`
	Quantity    uint   `json:"quantity"`

	gorm.Model
}

func (s *Quantity) BeforeCreate() error {
	s.UUID = uuid.New().String()
	return nil
}
