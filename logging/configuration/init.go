package configuration

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {

	var out io.Writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	var service string
	var level zerolog.Level = zerolog.TraceLevel
	var warning string
	if s, ok := os.LookupEnv("LOG_OUTPUT"); ok {
		switch s {
		case "JSON":
			out = os.Stdout
		default:
			warning = fmt.Sprintf("Unknown LOG_OUTPUT (%s); using console", s)
		}
	}
	if s, ok := os.LookupEnv("LOG_SERVICE_NAME"); ok {
		service = s
	} else {
		service = "eulabeia"
	}
	if s, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if l, err := zerolog.ParseLevel(s); err != nil {
			warning = fmt.Sprintf("Unable to identify log level (%s) fallback to TraceLevel", s)
		} else {
			level = l
		}
	}
	log.Logger = zerolog.New(out).With().Str("service", service).Caller().Timestamp().Logger().Level(level)
	if warning != "" {
		log.Warn().Msg(warning)
	}

}
