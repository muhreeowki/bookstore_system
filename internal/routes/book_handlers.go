package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muhreeowki/bookstore_system/internal/database"
	"gorm.io/gorm"
)

// Returns a single book by ID
func getBookHandler(c *fiber.Ctx) error {
	// Get the bookID from the URL
	param := struct {
		ID uint `params:"bookID"`
	}{}
	c.ParamsParser(&param)

	// Get the book by ID
	book, err := database.GetBookByID(database.Book{Model: gorm.Model{ID: param.ID}})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("error getting book: %v", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// Returns a list of all books
func getBooksHandler(c *fiber.Ctx) error {
	books, err := database.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

// Creates a new book and returns the created book
func createBookHandler(c *fiber.Ctx) error {
	var book database.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("error marshaling jSON: %v", err.Error()))
	}
	createdBook, err := database.CreateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error creating book: %v", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(createdBook)
}

// Deletes a book by ID
func deleteBookHandler(c *fiber.Ctx) error {
	// Get the bookID from the URL
	param := struct {
		ID uint `params:"bookID"`
	}{}
	c.ParamsParser(&param)

	// Get the book by ID
	err := database.DeleteBookByID(database.Book{Model: gorm.Model{ID: param.ID}})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("error getting book: %v", err.Error()))
	}

	return c.Status(fiber.StatusOK).SendString("Book deleted")
}
