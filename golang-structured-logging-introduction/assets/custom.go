package main

import (
	"context"
	"log/slog"
)

type CustomHandler struct {
	slog.Handler
}

func (c *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	// handle context here for tracing, monitoring, logging, etc
	return c.Handler.Handle(ctx, r)
}
