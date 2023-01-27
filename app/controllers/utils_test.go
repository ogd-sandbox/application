package controllers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func ServerSetup(t *testing.T, method, urlPath string, req *http.Request, handler func(f *fiber.Ctx) error) (int, []byte) {
	app := fiber.New()
	app.Add(method, urlPath, handler)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	return resp.StatusCode, body
}

func ControllerSetup(t *testing.T) *Controllers {
	controllers := NewControllers()

	return controllers
}
