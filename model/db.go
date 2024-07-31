package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/tweets?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Auto migrate your models
	db.AutoMigrate(&Tweet{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Like{})

	// Check if the users table exists
	if !db.Migrator().HasTable(&User{}) {
		log.Println("users table does not exist")
	} else {
		log.Println("users table exists")
	}

	// Check if a specific column exists in the users table
	if !db.Migrator().HasColumn(&User{}, "email") {
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

