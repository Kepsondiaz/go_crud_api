package main

import (
	"go_api_crud/initializer"

	"github.com/gin-gonic/gin"
)

func init() {

	initializer.LoadEnvVariable()
	
	
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue les gars",
		})
	})
	r.Run()
}