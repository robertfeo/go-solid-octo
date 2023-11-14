package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"backend/database"
	"backend/types"
)

var books = []types.Book{}

func getBooksHandler(c *gin.Context) {
	books, err := database.GetBooks() // Implement this function in database.go
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching books"})
		return
	}
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
	var newBook types.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func home(c *gin.Context) {
	c.String(http.StatusOK, "Hello Golang API")
}

func updateBookHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedBook types.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func initEnv() {
	godotenv.Load(".env")
}

func main() {
	initEnv()
	database.Init()
	router := gin.Default()
	router.StaticFile("/favicon.ico", "./favicon/favicon.ico")
	router.GET("/", home)
	router.GET("/books", getBooksHandler)
	router.GET("/books/:id", getBookByIDHandler)
	router.POST("/books", createBookHandler)
	router.PUT("/books/:id", updateBookHandler)
	router.DELETE("/books/:id", deleteBookHandler)
	router.Run(":8080")
}
