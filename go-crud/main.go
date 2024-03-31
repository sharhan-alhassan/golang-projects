package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sharhan-alhassan/golang-projects/go-crud/controllers"
	"github.com/sharhan-alhassan/golang-projects/go-crud/initializers"
)
 
func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main()  {
	r := gin.Default()

	// Create a single post instance
	r.POST("/post/create", controllers.PostCreate)

	// Retrieve all posts 
	r.GET("/post", controllers.PostRetrieveAll)

	// Retrieve a single post
	r.GET("/post/:id", controllers.PostRetrieve)

	// Update a post instance
	r.PUT("/post/:id", controllers.PostUpdate)

	// Delete post instance
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()	
}