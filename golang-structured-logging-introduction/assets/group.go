package main

import (
	"log/slog"
)

func foobar(logger *slog.Logger, args []string) {
	logger.Info("from foboar",
		slog.Any("args", args),
		slog.Bool("ok", true),
	)

}

func main() {
	foobarLogger := slog.
		With(slog.String("non_grouped", "value")).
		WithGroup("grouped")
	foobar(foobarLogger, []string{"a", "b"})
}
