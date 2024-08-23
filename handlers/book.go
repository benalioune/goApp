package handlers

import (
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all books
func GetBooks(c *gin.Context) {
	var books []models.Book
	// Logic to fetch books from DB will be added later
	c.JSON(http.StatusOK, books)
}

// PostBooks adds a new book
func PostBooks(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save book to DB will be added later
	c.JSON(http.StatusCreated, book)
}

// GetBookByID fetches a book by its ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	// Logic to fetch book by ID from DB will be added later
	c.JSON(http.StatusOK, book)
}
