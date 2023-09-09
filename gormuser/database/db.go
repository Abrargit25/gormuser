package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the database connection and returns a reference to the DB instance
func InitDB() *gorm.DB {
    // Replace the connection details with your actual database configuration
    dsn := ""
    
    // Initialize the database connection
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }
    
    // AutoMigrate your models here
    // Example: db.AutoMigrate(&User{}, &Product{})
    
    return db
}
