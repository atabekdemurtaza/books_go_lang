package unit

import (
	"testing"

	"book-management/models"
)

func TestBookModel(t *testing.T) {
	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
	}

	if book.Title != "Test Book" {
		t.Errorf("Expected Title to be 'Test Book', got '%s'", book.Title)
	}

	if book.Author != "Test Author" {
		t.Errorf("Expected Author to be 'Test Author', got '%s'", book.Author)
	}
}
