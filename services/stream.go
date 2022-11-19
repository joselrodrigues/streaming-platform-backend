package services

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type bufferBuilder interface {
	getAbsolutePath() error
	Open() error
	getSize() error
	setBuffer()
	getBuffer() *Video
}
type Video struct {
	file         *os.File
	buffer       []byte
	absolutePath string
	relativePath string
	size         int
}

type videoParams struct {
	id     string
	name   string
	source string
}

func (v *Video) getAbsolutePath() error {

	_, fileName, _, ok := runtime.Caller(0)

	if !ok {
		err := errors.New("it was not possible to recover the information")
		return err
	}

	basePath := filepath.Dir(fileName)
	v.absolutePath = filepath.Join(basePath, v.relativePath)

	return nil
}

func (v *Video) Open() error {

	file, err := os.Open(string(v.absolutePath))

	if err != nil {
		return err
	}

	v.file = file
	return err
}

func (v *Video) getSize() error {

	fileInfo, err := v.file.Stat()

	if err != nil {
		return err
	}
	v.size = int(fileInfo.Size())
	return nil
}

func (v *Video) setBuffer() {
	buffer := make([]byte, v.size)
	v.file.Read(buffer)
	v.buffer = buffer
}

func (v *Video) getBuffer() *Video {
	return v
}

func buildBuffer(b bufferBuilder) (*Video, error) {
	err := b.getAbsolutePath()
	if err != nil {
		return nil, err
	}
	err = b.Open()
	if err != nil {
		return nil, err
	}
	err = b.getSize()
	if err != nil {
		return nil, err
	}
	b.setBuffer()
	return b.getBuffer(), nil
}

func setBuilderBuffer(t *videoParams) *Video {
	video := &Video{}
	relativePath := fmt.Sprintf("../videos/%s/%s/%s.mp4", t.source, t.name, t.id)
	video.relativePath = relativePath

	return video
}

func StreamService(c *fiber.Ctx) error {
	params := &videoParams{c.Params("id"), c.Params("name"), c.Params("source")}
	builder := setBuilderBuffer(params)
	video, err := buildBuffer(builder)

	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	defer video.file.Close()

	c.Set("Content-Length", fmt.Sprintf("%d", video.size))
	c.Set("Content-Type", "video/mp4")
	c.Set("Connection", "keep-alive")
	c.Set("Accept-Ranges", "bytes")
	c.Set("Content-Range", fmt.Sprintf("bytes 0-%[1]d/%[1]d", video.size))
	c.Response().SetBodyStream(bytes.NewReader(video.buffer), video.size)

	return nil
}
