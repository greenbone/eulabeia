package logging

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var out io.Writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

var service string = "eulabeia"

var level zerolog.Level = zerolog.TraceLevel

func SetOutput(writer io.Writer) {
	out = writer
}

func SetService(s string) {
	service = s
}

func SetLevel(l zerolog.Level) {
	level = l
}

func Logger() zerolog.Logger {
	return zerolog.New(out).With().Str("service", service).Caller().Timestamp().Logger().Level(level)
}
