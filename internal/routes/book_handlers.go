package routes

import (
	"github.com/gofiber/fiber/v2"
)

func getBook(c *fiber.Ctx) error {
	return c.SendString("Get Book")
}

func getAllBooks(c *fiber.Ctx) error {
	return c.SendString("Get All Books")
}

func createBook(c *fiber.Ctx) error {
	return c.SendString("Add A Book")
}

func updateBook(c *fiber.Ctx) error {
	return c.SendString("Update a Book")
}

func deleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete a Book")
}
