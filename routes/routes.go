package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()

	StreamRoutes(app)
	Swagger(app)

	return app
}
