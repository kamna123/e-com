package models

type Address struct {
	UUID   string `json:"uuid" gorm:"unique;not null;index;primaryKey"`
	UserID string `json:"user_id"`

	StreetAddress string `gorm:"not null"`
	City          string `gorm:"not null"`
	Country       string `gorm:"not null"`
	ZipCode       string `gorm:"not null"`
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	PhoneNumber   string `gorm:"not null"`
}
