package initializer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
  DB := *gorm.db

func ConnectToDB() {
	err := Error
	dsn := "user:pass@tcp(127.0.0.1:3306)/go_api_?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}