package app

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Int("pid", os.Getpid()).
	Logger()

func GetLogger() *zerolog.Logger {
	return &logger
}

func LogFatal(err error) {
	logger.Fatal().Msg(err.Error())
}
