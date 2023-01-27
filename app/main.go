package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/nobuyo/nrfiber"

	"github.com/ogd-sanbox/application/pkg/configs"
	"github.com/ogd-sanbox/application/pkg/middleware"
	"github.com/ogd-sanbox/application/pkg/routes"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Add New Relic.
	nrapp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("application"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)
	if err != nil {
		log.Printf("Error while connecting New Relic: %v", err)
	}

	app.Use(nrfiber.New(nrfiber.Config{
		NewRelicApp: nrapp,
	}))

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PublicRoutes(app) // Register a public routes for app.

	// Start server.
	app.Listen(":3333")
}
