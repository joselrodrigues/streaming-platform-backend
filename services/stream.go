package services

import (
	"bytes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func StreamService(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.Params("name")
	source := c.Params("source")
	video := fmt.Sprintf("videos/%s/%s/%s.mp4", source, name, id)

	file, err := os.Open(string(video))
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	fileSize := int(fi.Size())

	buffer := make([]byte, fileSize)
	file.Read(buffer)

	c.Set("Content-Length", fmt.Sprintf("%d", fileSize))
	c.Set("Content-Type", "video/mp4")
	c.Set("Connection", "keep-alive")
	c.Set("Accept-Ranges", "bytes")
	c.Set("Content-Range", fmt.Sprintf("bytes 0-%[1]d/%[1]d", fileSize))
	c.Response().SetBodyStream(bytes.NewReader(buffer), fileSize)

	return nil
}
