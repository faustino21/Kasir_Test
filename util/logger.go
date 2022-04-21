package util

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func NewLog(loglevel string) {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if loglevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = logger
}
