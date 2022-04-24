package util

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strconv"
	"time"
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

func TimeUnix(unix int) *time.Time {
	i, err := strconv.ParseInt(fmt.Sprintf("%d", unix), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return &tm
}
