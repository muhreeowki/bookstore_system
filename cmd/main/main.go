package main

import (
	"fmt"
	"log"

	"github.com/muhreeowki/bookstore_system/internal/config"
	"github.com/muhreeowki/bookstore_system/internal/routes"
)

func main() {
	// Connect to the database
	db, err := config.StartDB()
	if err != nil {
		panic(err)
	}
	log.Println("DB connection successful", db)

	// Setup the Router
	app := routes.NewRouter()

	// Setup url
	port := 8000
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server is running on port %d\n", port)
	app.Listen(addr)
}
