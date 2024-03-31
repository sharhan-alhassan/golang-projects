package main

import (
	"github.com/sharhan-alhassan/golang-projects/go-crud/initializers"
	"github.com/sharhan-alhassan/golang-projects/go-crud/models"
)


func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	// // Migrate the Schema
	// initializers.DB.AutoMigrate(&models.Post{})
}

func main() {
	// Migrate the Schema
	initializers.DB.AutoMigrate(&models.Post{})
}