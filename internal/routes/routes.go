package routes

import "github.com/gofiber/fiber/v2"

func NewRouter() *fiber.App {
	// Create a new router
	app := fiber.New()

	// Setup Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Get("/book", getBookHandler)
	app.Get("/books", getBooksHandler)
	app.Post("/book", createBookHandler)
	app.Put("/book", updateBookHandler)
	app.Delete("/book", deleteBookHandler)

	return app
}
