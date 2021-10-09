package models

import "github.com/google/uuid"

type Cart struct {
	UUID        string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	UserID      string `json:"userid"`
	ProductUUID string `json:"product_uuid" gorm:"not null;index"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
}

func (s *Cart) BeforeCreate() error {
	s.UUID = uuid.New().String()
	return nil
}
