package database

import (
	"fmt"

	"log"

	"github.com/jinzhu/gorm"

	"e-commerce/cmd/config"
)

var Database *gorm.DB

func init() {
	dbConfig := config.Config.Database
	connectionPath := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password, dbConfig.SSLMode)

	database, err := gorm.Open("postgres", connectionPath)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	Database = database
}
