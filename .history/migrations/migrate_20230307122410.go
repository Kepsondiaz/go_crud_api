package main

import "go_api_crud/initializer"
import "go_api_crud/"

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.posts{})
}
