package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buraksaglam089/tool-tube/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))

	log.Printf("Attempting to connect to database with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(types.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate db : %v", err)
	}

	log.Println("Successfully connected to database")

	return db, nil
}
