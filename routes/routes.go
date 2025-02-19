package routes

import (
	"github.com/dancankarani/delha-frontend/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	
	api := app.Group("/api/v1")

	api.Post("/listings", controller.CreateListing)
	api.Get("/listings", controller.GetListings)
	api.Get("/listings/:id", controller.GetListing)
	api.Patch("/listings/:id", controller.UpdateListing)
	api.Delete("/listings/:id", controller.DeleteListing)
}

func SetAuthRoutes(app *fiber.App){
	api := app.Group("/api/v1")
	api.Post("/auth/admin", controller.AdminLogin)
}