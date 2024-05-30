package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Configer DB configuration
func StartDB() (*gorm.DB, error) {
	// Connection string
	dsn := "host=ep-misty-pond-a2qbqwcd.eu-central-1.aws.neon.tech user=didomi_owner password=RWfJu0Zj8bEB dbname=bookstore_system sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
