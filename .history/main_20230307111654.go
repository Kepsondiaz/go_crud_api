package main 



import (
	"github.com/gin-gonic/gin"

	"log"

    "github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	r := gin.Default()
}

func main() {
	

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue les gars",
		})
	})
	r.Run()
}