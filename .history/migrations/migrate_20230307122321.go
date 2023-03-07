package main

import "go_api_crud/initializer"

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&)
}
