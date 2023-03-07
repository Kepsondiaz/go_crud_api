package main 



import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue les gars",
		})
	})
	r.Run()
}