package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Configer DB configuration
func StartDB() error {
	// Connection string
	dsn := "host=ep-misty-pond-a2qbqwcd.eu-central-1.aws.neon.tech user=didomi_owner password=RWfJu0Zj8bEB dbname=bookstore_system sslmode=require"

	// Open the connection and store it in the global DB variable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	// Migrate the schema
	DB.AutoMigrate(&Book{})

	// Add Books

	log.Println("DB connection successful", db)
	return nil
}
