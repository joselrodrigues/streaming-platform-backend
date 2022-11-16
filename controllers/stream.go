package controllers

import (
	s "stream/services"

	"github.com/gofiber/fiber/v2"
)

// StreamController godoc
// @Summary      Show an account
// @Description  get the video by source, name and id
// @Tags         stream
// @Produce		 octet-stream
// @Param        source   path      string  true  "Video Source"
// @Param        name     path      string  true  "Video Name"
// @Param        id       path      int     true  "Video Number"
// @Success      200
// @Failure      500
// @Router       /{source}/{name}/{id} [get]
func StreamController(c *fiber.Ctx) error {
	err := s.StreamService(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}
