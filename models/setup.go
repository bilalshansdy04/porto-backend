package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// DSN (Data Source Name)
	// Format: username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:@tcp(127.0.0.1:3308)/porto?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! Please make sure database 'porto' exists.", err)
	}

	err = database.AutoMigrate(&Profile{}, &Project{}, &Experience{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	DB = database
}
