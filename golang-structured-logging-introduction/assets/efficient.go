package main

import (
	"context"
	"log/slog"
	"time"
)

func main() {
	slog.LogAttrs(context.Background(), // context
		slog.LevelInfo,          // level
		"your message",          // message
		slog.Float64("key", 10), // keys
		slog.Group("latency", // grouping keys
			slog.Duration("duration", time.Millisecond),
			slog.Time("start", time.Now()),
		),
	)
}
