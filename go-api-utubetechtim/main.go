package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "In search of the lost ship", Author: "Joe Doe", Quantity: 3},
	{ID: "2", Title: "The count of monte cristo", Author: "Sam Jame", Quantity: 5},
	{ID: "3", Title: "Fourty nights of faith", Author: "Sharhan Alhassan", Quantity: 2},
}

// Return the JSON version of all the books
func getBooks(c *gin.Context) {

	// Nicely format data into JSON
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Books retrieved successfully",
		"data":    books,
		"status":  http.StatusOK,
	})
}

// call the getBookById function
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"data": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}


// Checkout book -- Using query parameter
func checkoutBook(c *gin.Context)  {
	// ok parameter returns true/false if "id" in GetQuery param is returned or not
	id, ok:= c.GetQuery("id")
	
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"data": "Missing query parameter",
		})
		return
	}

	// Get the book by "Id"
	book, err := getBookById(id)

	// If error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"data": "book not found",
		})
		return
	}

	// If book quantity is less than/equal to zero
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"data": "book not available",
		})
	}

	// If everything fine, increment book Quantity
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book checked out",
		"data": book,
	})
}


// Function to Get book by id
func getBookById(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// Add post to the slice
func createBook(c *gin.Context) {
	// Add book
	var newBook book // Is of type book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Equivalent to
	// err := c.BindJSON(&newBook)
	// if err != nil {
	// 	return
	// }

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	// set up router
	router := gin.Default()

	// GET request to list all books
	router.GET("/books", getBooks)

	// POST request to add a book to the slice
	router.POST("/books", createBook)

	// Get request for a single book
	router.GET("/books/:id", bookById)

	// Patch request to check out book
	router.PATCH("/checkout", checkoutBook)
	// Run server
	router.Run("localhost:8080")

}
