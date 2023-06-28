package pkg

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"os"
)

// LogFromCtx returns a Logger from a context. If no logger is found in the context, a new one is created using a default text handler.
func LogFromCtx(ctx context.Context) *slog.Logger {
	v, ok := ctx.Value("logger").(*slog.Logger)
	if ok {
		return v
	}
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

// LogIntoCtx returns a new context with the given logger.
func LogIntoCtx(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

func userFromEchoContext(c echo.Context) string {
	email := ""
	if v, ok := c.Get("email").(string); ok {
		email = v
	}
	return email
}
