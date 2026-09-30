package main

import (
	"fmt"
	"log"

	"backend/database"
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Inisialisasi Database
	database.ConnectDB()

	// Setup Fiber App
	app := fiber.New()
	app.Use(cors.New())

	// Routing
	app.Post("/api/register", handlers.Register)

	// Jalankan Server
	fmt.Println("Server Golang Fiber berjalan di http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}