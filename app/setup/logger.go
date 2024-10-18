package setup

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func SetupLog() *slog.Logger {

	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})

	Logger = slog.New(h)
	slog.SetDefault(Logger)

	programLevel.Set(slog.LevelDebug)

	return Logger
}
