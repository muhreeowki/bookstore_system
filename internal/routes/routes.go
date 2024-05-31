package routes

import "github.com/gofiber/fiber/v2"

func NewRouter() *fiber.App {
	// Create a new router
	app := fiber.New()

	// Setup Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("/book", createBookHandler)
	app.Get("/book/:bookID", getBookHandler)
	app.Delete("/book/:bookID", deleteBookHandler)
	app.Get("/books", getBooksHandler)

	return app
}
