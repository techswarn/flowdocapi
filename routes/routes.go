package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/flowdocapi/handlers"
	"github.com/techswarn/flowdocapi/middleware"
)

// SetupRoutes register routes based on functionalities
func SetupRoutes(app *fiber.App) {
	// public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")

	//Health checks and other data
	publicRoutes.Get("/", handlers.Addnodehandler)


}