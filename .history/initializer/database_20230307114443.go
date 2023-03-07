package initializer

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  var DB *gorm.DB

func ConnectToDB() {
	var err error
	con := "root:@tcp(127.0.0.1:3306)/go_api_crud?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(con), &gorm.Config{})

	if err != nil {
		log.Fatal("fail to connect to database")
	} 
}

