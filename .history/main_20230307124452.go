package main

import (
	"go_api_crud/initializer"

	"github.com/gin-gonic/gin"

	"go_api_crud/controllers"
)

func init() {

	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
	
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.createPosts())
	r.Run()
}