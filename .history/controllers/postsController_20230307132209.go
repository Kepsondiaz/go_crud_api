package controllers

import (
	"github.com/gin-gonic/gin"

	"go_api_crud/models"
)

func CreatePosts(c *gin.Context) {

	// create data 
	posts := models.Posts{Title: "Saint Valentin", content: "le mois de l'amour"}

	c.JSON(200, gin.H{
		"message": "Bienvenue les gars",
	})
}