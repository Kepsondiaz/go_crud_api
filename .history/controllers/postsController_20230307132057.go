package controllers

import (
	"github.com/gin-gonic/gin"

	"go_api_crud/models"
)

func CreatePosts(c *gin.Context) {

	// create data 
	posts := Posts{Title: "Saint Valentin", Age: 18, Birthday: time.Now()}

	c.JSON(200, gin.H{
		"message": "Bienvenue les gars",
	})
}