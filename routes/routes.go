package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/flowdocapi/handlers"
)

// SetupRoutes register routes based on functionalities
func SetupRoutes(app *fiber.App) {
	// public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")

	//Health checks and other data
	publicRoutes.Post("/node", handlers.Addnodehandler)
	publicRoutes.Get("/node", handlers.GetNodeshandler)
	publicRoutes.Get("/edge", handlers.GetEdgeshandler)
	publicRoutes.Get("/article", handlers.GetArticle)
}