package main

import (
	"log/slog"
	"os"
)

var (
	level  = new(slog.LevelVar)
	logger = slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{Level: level}))
)

func main() {
	logger.Debug("won't be shown") // by default log level is INFO
	level.Set(slog.LevelDebug)     // change it to DEBUG
	logger.Debug("shown")
}
