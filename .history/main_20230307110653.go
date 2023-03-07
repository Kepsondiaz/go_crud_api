package main 

import "fmt"

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue les gars",
		})
	})
	r.Run()
}