package main

import (
	logger "stream/log"
	routes "stream/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()
	logger.Setup()
	routes.Setup(app)
	log.Info().Msg("Server is listening on port 3000")
	app.Listen(":3000")
}
