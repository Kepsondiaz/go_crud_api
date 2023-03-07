package main

import (
	"go_api_crud/initializer"

	"go_api_crud/inimodels"
)

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.posts{})
}
