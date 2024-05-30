package routes

import "github.com/gofiber/fiber/v2"

func NewRouter() *fiber.App {
	// Create a new router
	app := fiber.New()

	// Setup Routes
	app.Get("/book", getBook)
	app.Get("/books", getAllBooks)
	app.Post("/book", createBook)
	app.Put("/book", updateBook)
	app.Delete("/book", deleteBook)

	return app
}
