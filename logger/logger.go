package logger

import (
	"os"
	"sync"

	"golang.org/x/exp/slog"
)

var logger *slog.Logger
var once sync.Once

func getOrCreateLogger() *slog.Logger {
	once.Do(func() {
		logFile, err := os.OpenFile("storage/go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			println(err.Error())
		}

		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler := slog.NewJSONHandler(logFile, opts)
		logger = slog.New(handler)
	})
	return logger
}

func NewLogger() *slog.Logger {
	return getOrCreateLogger()
}
