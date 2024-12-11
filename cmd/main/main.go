package main

import (
	"fmt"
	"log"

	"book-management/handlers"
	"book-management/models"
	"book-management/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database setup
func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// Auto-migrate the Book model
	db.AutoMigrate(&models.Book{})
	return db
}

// Gin router setup
func NewRouter(h *handlers.BookHandler) *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r, h)
	return r
}

// Main function to start the application
func main() {
	app := fx.New(
		fx.Provide(NewDB),
		fx.Provide(handlers.NewBookHandler),
		fx.Provide(NewRouter),
		fx.Invoke(func(r *gin.Engine) {
			fmt.Println("Server running at http://localhost:8080")
			r.Run(":8080") // Start the server
		}),
	)

	app.Run()
}
