package controllers

import "github.com/gin-gonic/gin"

func createPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Bienvenue les gars",
	})
}