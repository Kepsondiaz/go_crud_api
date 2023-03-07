package controllers

import "github.com/gin-gonic/gin"

func CreatePosts(c *gin.Context) {

	// create data 
	posts := Posts{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	c.JSON(200, gin.H{
		"message": "Bienvenue les gars",
	})
}