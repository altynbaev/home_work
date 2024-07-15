package observability

import (
	"context"
	"log/slog"
)

func DefaultLogger() slog.Logger {

}

// CtxLogger возвращает логгер, внедренный в ctx. В противном случае
// возвращается дефолтный логгер проекта.
func CtxLogger(ctx context.Context) slog.Logger {
	if logger, ok := ctx.Value("loggerKey").(slog.Logger); ok {
		return logger
	}
	return DefaultLogger()
}

func CtxWithLogger(ctx context.Context, lg slog.Logger) context.Context {
	return context.WithValue(ctx, "logger", lg)
}
