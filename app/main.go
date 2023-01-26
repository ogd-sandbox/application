package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ogd-sanbox/application/pkg/configs"
	"github.com/ogd-sanbox/application/pkg/middleware"
	"github.com/ogd-sanbox/application/pkg/routes"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PublicRoutes(app) // Register a public routes for app.

	// Start server.
	app.Listen(":3333")
}
