package initializer

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  var DB *gorm.DB

func ConnectToDB() {
	var err error
	con := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(con), &gorm.Config{})

	if err != nil {
		log.Fatal("fail to connect to database")
	} 
}

