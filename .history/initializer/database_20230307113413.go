package initializer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
  DB := *gorm.db

func ConnectToDB() {
	err := Error
	dsn := "user:@tcp(127.0.0.1:3306)/go_api_crud?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}