package controllers

import (
	"github.com/gin-gonic/gin"

	"go_api_crud/initializer"
	"go_api_crud/models"
)

// create the the posts
func CreatePosts(c *gin.Context) {

	var bodyPost struct {
		Title string
		Content string
	}
	c.Bind(&bodyPost)
	// create data 
	posts := models.Posts{Title: bodyPost.Title, Content: bodyPost.Content}

	result := initializer.DB.Create(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// show the posts
func ShowPosts(c *gin.Context) {
 	var posts []models.Posts
 	initializer.DB.Find(&posts)

	 c.JSON(200, gin.H{
		"posts": posts,
	})

 }

//update the posts 
// func UpdatePosts(c *gin.Context) {
// 	initializer.DB.Where("id = ?", 1).Update("Title", "hello")

// 	c.JSON(200, gin.H{
// 		"posts": posts,
// 	})
// }