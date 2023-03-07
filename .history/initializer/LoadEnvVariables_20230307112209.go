package initializer  

import (

	"log"

    "github.com/joho/godotenv"
)

func LoadEnvVariable()

	func init() {
		err := godotenv.Load()
	
		if err != nil {
		  log.Fatal("Error loading .env file")
		}

}

