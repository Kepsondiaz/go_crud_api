package initializer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
  vDB := *gorm.db

func ConnectToDB() {
	err := Error
	dsn := "root:@tcp(127.0.0.1:3306)/go_api_crud?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}