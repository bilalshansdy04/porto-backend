package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load, will use system environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")

	// DSN (Data Source Name)
	// Format: username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local&tls=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		username, password, host, port, databaseName)
	
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! Error: ", err)
	}

	err = database.AutoMigrate(&Profile{}, &Project{}, &Experience{}, &Skill{}, &Setting{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	DB = database
}
