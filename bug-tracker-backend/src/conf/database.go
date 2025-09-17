package conf

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	env := os.Getenv("ENVIRONMENT")

	// Check if required environment variables are set
	if db_host == "" || db_user == "" || db_password == "" || db_port == "" || db_name == "" {
		log.Printf("Warning: Database environment variables not fully configured")
		log.Printf("DB_HOST: %s, DB_USERNAME: %s, DB_PORT: %s, DB_NAME: %s", db_host, db_user, db_port, db_name)
		return
	}

	ssl_mode := ""
	if env != "local" {
		ssl_mode = "require"
	} else {
		ssl_mode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC", db_host, db_user, db_password, db_name, db_port, ssl_mode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Warning: Failed to connect to database: %v", err)
		log.Printf("Application will continue but database operations will fail")
		return
	}

	log.Printf("Successfully connected to database")
}
