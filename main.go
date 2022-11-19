package main

import (
	_ "stream/docs"
	logger "stream/log"
	routes "stream/routes"

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
// @host localhost:4000
// @BasePath /
func main() {
	logger.Setup()

	app := routes.Setup()

	err := app.Listen(":4000")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
