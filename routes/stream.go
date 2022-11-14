package routes

import (
	c "stream/controllers"

	"github.com/gofiber/fiber/v2"
)

func StreamRoutes(app *fiber.App) {
	app.Get("/:source/:name/:id", c.StreamController)
}
