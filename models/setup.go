package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load, will use system environment variables")
	}

	tursoURL := os.Getenv("TURSO_DATABASE_URL")
	tursoToken := os.Getenv("TURSO_AUTH_TOKEN")
	var database *gorm.DB

	if tursoURL != "" {
		dsn := tursoURL
		if tursoToken != "" {
			dsn += "?authToken=" + tursoToken
		}
		
		sqlDB, err := sql.Open("libsql", dsn)
		if err != nil {
			log.Fatal("Failed to open libSQL connection:", err)
		}

		database, err = gorm.Open(sqlite.New(sqlite.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

		if err != nil {
			log.Fatal("Failed to connect to Turso database:", err)
		}
		log.Println("Connected to Turso database!")
	} else {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		databaseName := os.Getenv("DB_DATABASE")

		// DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
			username, password, host, port, databaseName)
		
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("Failed to connect to MySQL/TiDB database! Error: ", err)
		}
		log.Println("Connected to MySQL/TiDB database!")
	}

	err = database.AutoMigrate(&Profile{}, &Project{}, &ImageScreenshot{}, &Experience{}, &Skill{}, &Setting{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	DB = database
}
