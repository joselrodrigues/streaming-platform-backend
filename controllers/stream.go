package controllers

import (
	s "stream/services"

	"github.com/gofiber/fiber/v2"
)

func StreamController(c *fiber.Ctx) error {
	return s.StreamService(c)
}
