package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ogd-sanbox/application/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create controllers.
	c := controllers.NewControllers()

	// Create routes group.
	v1 := a.Group("/api/v1")

	// Routes for GET method:
	v1.Get("/health", c.Health()) // Health check.
}
