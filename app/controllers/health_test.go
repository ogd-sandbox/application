package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
)

func TestControllers_Health(t *testing.T) {
	controller := ControllerSetup(t)
	req := httptest.NewRequest("GET", "/api/v1/health", nil)

	statusCode, _ := ServerSetup(
		t, "GET", "/api/v1/health",
		req, controller.Health(),
	)
	assert.Equal(t, fiber.StatusOK, statusCode)
}
