package controllers

import "github.com/gofiber/fiber/v2"

// Health returns ok to show the application is up and running
func (c *Controllers) Health() fiber.Handler {
	return func(f *fiber.Ctx) error {
		return f.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": "OK",
			},
		)
	}
}
