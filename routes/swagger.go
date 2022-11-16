package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Swagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
