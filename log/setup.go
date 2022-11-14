package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup() {
	output := zerolog.ConsoleWriter{Out: os.Stderr}

	output.FormatMessage = func(i interface{}) string {
		if i != nil {
			return fmt.Sprintf("***%s****", i)
		}
		return ""
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	log.Logger = log.Output(output)
}
