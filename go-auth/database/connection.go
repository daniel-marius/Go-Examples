package database

import (
  "go-auth/models"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

// Global variable
var DB *gorm.DB

func Connect() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  endpointURL := "root:@tcp(127.0.0.1:3306)/go-auth?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(endpointURL), &gorm.Config{})

  if err != nil {
    panic("Could not connect to the database!")
  }

  DB = db

  // Creates the table inside db
  db.AutoMigrate(&models.User{})
}
