package main

import (
	"log/slog"
	"net/http"
	"os"
)

var (
	textLogger *slog.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	jsonLogger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
)

func main() {
	slog.Info("this is info using default logger")
	textLogger.Warn("this is warning", slog.Int("response_code", http.StatusNotFound))
	jsonLogger.Error("this is error",
		slog.String("status", http.StatusText(http.StatusInternalServerError)),
		slog.Bool("is_json", true),
	)
}
