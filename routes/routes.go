package routes

import (
	"book-management/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handlers.BookHandler) {
	r.GET("/books", h.GetAllBooks)
	r.GET("/books/:id", h.GetBookByID)
	r.POST("/books", h.CreateBook)
	r.DELETE("/books/:id", h.DeleteBook)
}
