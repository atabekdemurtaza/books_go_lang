package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Database setup
func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// Auto-migrate the Book model
	db.AutoMigrate(&Book{})
	return db
}

// Gin router setup
func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Route: Get all books
	r.GET("/books", func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		c.JSON(http.StatusOK, books)
	})

	// Route: Get a book by ID
	r.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		var book Book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	})

	// Route: Create a new book
	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&book)
		c.JSON(http.StatusCreated, book)
	})

	// Route: Delete a book
	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Book{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	})

	return r
}

// Main function to start the application
func main() {
	app := fx.New(
		fx.Provide(NewDB),
		fx.Provide(NewRouter),
		fx.Invoke(func(r *gin.Engine) {
			fmt.Println("Server running at http://localhost:8080")
			r.Run(":8080") // Start the server
		}),
	)

	app.Run()
}
