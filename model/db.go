package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global variable for the database connection
var db *gorm.DB

// InitializeDB initializes the database connection.
// Connects to a MySQL database using the root user at localhost:3306.
// Creates the tweets, users, and likes tables if they do not exist.
// Checks if the users table exists and if the email column exists in the users table.
// InitializeDB initializes the database connection
func InitializeDB() {
	// connect with mysql of 3306 port
	var err error
	dsn := os.Getenv("DB_DSN")
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Create tables if they don't exist
	db.AutoMigrate(&User{}, &Tweet{}, &Like{})
	// Auto migrate your models
	db.AutoMigrate(&Tweet{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Like{})
}

// GetDB returns the database connection.
// Returns an error if the database connection has not been initialized.
// GetDB returns the database connection
func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return db, nil
}


