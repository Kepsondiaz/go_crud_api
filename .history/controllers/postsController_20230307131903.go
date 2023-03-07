package controllers

import "github.com/gin-gonic/gin"

func CreatePosts(c *gin.Context) {

	// create data 
	

	c.JSON(200, gin.H{
		"message": "Bienvenue les gars",
	})
}