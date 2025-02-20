package main

import (
	"fmt"

	"github.com/dancankarani/delha-frontend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("hello, world")

	// Setup routes
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://delha-frontend.vercel.app/", // Replace with your frontend URL
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true, // This must be true for cookies or authentication
	}))
	
	routes.SetupRoutes(app)
	routes.SetAuthRoutes(app)

	// Start server
	app.Listen(":8000")
}