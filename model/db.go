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
	var err error
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

	// Check if the users table exists
	if !db.Migrator().HasTable(&User{}) {
		log.Fatal("users table does not exist")
		log.Println("users table does not exist")
	} else {
		log.Println("users table exists")
	}

	// Check if the email column exists in the users table
	// Check if a specific column exists in the users table
	if !db.Migrator().HasColumn(&User{}, "email") {
		log.Fatal("email column does not exist in users table")
		log.Println("email column does not exist in users table")
	} else {
		log.Println("email column exists in users table")
	}
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


