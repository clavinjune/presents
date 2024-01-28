package main

import (
	"context"
	"log/slog"
	"time"
)

func main() {
	slog.Info("❌ have more allocations since we don't define the type",
		"key", "value",
		"key2", time.Minute,
	)
	slog.LogAttrs(context.Background(), slog.LevelInfo,
		"✅ avoid extra allocations",
		slog.String("key", "value"),
		slog.Duration("key2", time.Minute),
	)
}
