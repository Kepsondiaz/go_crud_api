package initializer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )

func ConnectToDB() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB,  := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}