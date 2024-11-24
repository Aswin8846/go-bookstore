package config

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
  // Adjust the DSN with your credentials (e.g., user:password@tcp(localhost:3306)/dbname)
  d, err := gorm.Open(mysql.Open("root:haywin$8846@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
  if err != nil {
    panic("failed to connect to the database")
  }
  db = d
}

func GetDB() *gorm.DB {
  return db
}

