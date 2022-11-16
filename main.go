package main

import (
	_ "stream/docs"
	logger "stream/log"
	routes "stream/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// @title Stream API
// @version 1.0
// @description Stream API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /
func main() {
	app := fiber.New()
	logger.Setup()
	routes.Setup(app)

	err := app.Listen(":4000")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
