package main

import (
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{ReplaceAttr: func(g []string, a slog.Attr) slog.Attr {
			if "password" == a.Key {
				a.Value = slog.StringValue("[PII DATA]")
			}
			return a
		}}))
)

func main() {
	logger.Info("password will be redacted",
		slog.String("password", "P@$$w0rd?!"))
}
