package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 5},
}

func getBooksHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByIDHandler(c *gin.Context) {
	// path parameter
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBookHandler(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooksHandler)
	router.GET("/books/:id", getBookByIDHandler)
	router.POST("/books", createBookHandler)
	/* router.PUT("/books/:id", updateBookHandler)
	router.DELETE("/books/:id", deleteBookHandler) */
	router.Run(":8080")
}
