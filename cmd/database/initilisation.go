package database

import (
	"fmt"
	"log"

	"github.com/golang/glog"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"

	"e-commerce/cmd/config"
)

var Database *gorm.DB

func init() {
	dbConfig := config.Config.Database
	connectionPath := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password)
	glog.Info("database config - %s", connectionPath)
	database, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ecom?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	Database = database
}
