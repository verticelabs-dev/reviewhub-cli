package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Int("pid", os.Getpid()).
	Logger()

func GetLogger() *zerolog.Logger {
	return &Logger
}
