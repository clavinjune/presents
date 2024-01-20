package main

import (
	"log/slog"
)

func main() {
	logger2 := slog.With(slog.Int("num", 3))

	logger2.Info("whenever logger2 is called, you can see the num attribute derived")
}
