package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharhan-alhassan/golang-projects/go-crud/initializers"
	"github.com/sharhan-alhassan/golang-projects/go-crud/models"
)

// Post create a single post instance
func PostCreate(c *gin.Context) {
	// Fetch data from request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return the post

	c.JSON(200, gin.H{
		"data": post,
	})
}

// Post read all posts instances
func PostRetrieveAll(c *gin.Context) {
	// Get all posts
	// create posts variable as an array
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	// If there's error in retrieving all posts, throw it
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return all posts to API
	c.JSON(200, gin.H{
		"data": posts,
	})
}

// Retrieve a single user by ID
func PostRetrieve(c *gin.Context) {
	// Fetch id from URL params
	id := c.Param("id")

	// Get the post
	var post models.Post

	// Filter the post by the "ID"
	result := initializers.DB.First(&post, id)

	// Check if an error occurred during the database operation
	if result.Error != nil {
		// Handle the error
		log.Println("Error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	// Check if no rows were affected, meaning no post found
	if result.RowsAffected == 0 {
		// Return a custom error message and status code
		log.Println("Post not found")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	// No error and post found, return the response
	c.JSON(200, gin.H{
		"data": post,
	})
}

// 
func PostUpdate(c *gin.Context) {
	// Get the id from the request body
	id := c.Param("id")

	// Get the data from the request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post we're updating
	var post models.Post
	// result := initializers.DB.First(&post, id)
	initializers.DB.First(&post, id)

	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// 	return
	// }

	// Update the instance
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Return the response data
	c.JSON(200, gin.H{
		"data": post,
	})
}


func PostDelete(c *gin.Context)  {
	// Get post id from requestl URL
	id := c.Param("id")

	// Find the post object we want to delete
	var post models.Post
	initializers.DB.Delete(&post, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}