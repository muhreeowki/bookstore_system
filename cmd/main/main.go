package main

import (
	"fmt"

	"github.com/muhreeowki/bookstore_system/internal/database"
	"github.com/muhreeowki/bookstore_system/internal/routes"
)

func main() {
	// Setup the Router
	app := routes.NewRouter()

	// Setup url
	port := 8000
	addr := fmt.Sprintf(":%d", port)

	// Start db connection
	err := database.StartDB()
	if err != nil {
		panic(err)
	}

	// Start the server
	fmt.Printf("Server is running on port %d\n", port)
	err = app.Listen(addr)
	if err != nil {
		panic(err)
	}
}
