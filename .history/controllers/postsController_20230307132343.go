package controllers

import (
	"github.com/gin-gonic/gin"

	"go_api_crud/initializer"
	"go_api_crud/models"
)

func CreatePosts(c *gin.Context) {

	// create data 
	posts := models.Posts{Title: "Saint Valentin", Content: "le mois de l'amour"}

	result = initializer.DB.Create(&posts)

	if result != nil {
		
	}

	c.JSON(200, gin.H{
		"message": "Bienvenue les gars",
	})
}