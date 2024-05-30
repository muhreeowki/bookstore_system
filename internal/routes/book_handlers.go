package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muhreeowki/bookstore_system/internal/database"
)

func getBookHandler(c *fiber.Ctx) error {
	return c.SendString("Get Book")
}

func getBooksHandler(c *fiber.Ctx) error {
	books, err := database.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func createBookHandler(c *fiber.Ctx) error {
	var book database.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error Marshaling JSON: %v", err.Error()))
	}
	createdBook, err := database.CreateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error Creating Book: %v", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(createdBook)
}

func updateBookHandler(c *fiber.Ctx) error {
	return c.SendString("Update a Book")
}

func deleteBookHandler(c *fiber.Ctx) error {
	return c.SendString("Delete a Book")
}
