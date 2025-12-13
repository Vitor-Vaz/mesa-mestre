package telemetryfs

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// loggerKey holds the context key for the logger.
type loggerKey struct{}

// WithLogger adds the provided zap.Logger to the context.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// NewLogger creates and returns a new zap.Logger instance configured for use.
func NewLogger() (*zap.Logger, error) {
	logConfig := zap.NewProductionConfig()
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)
	return logger, nil
}

func fallbackLogger() *zap.Logger {
	logger := zap.L()
	logger.Error("missing logger on context")

	return logger
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return fallbackLogger()
	}

	logger, ok := ctx.Value(loggerKey{}).(*zap.Logger)
	if !ok || logger == nil {
		return fallbackLogger()
	}

	return logger
}

func Info(ctx context.Context, msg string) {
	LoggerFromContext(ctx).Info(msg)
}

func Error(ctx context.Context, msg string) {
	LoggerFromContext(ctx).Error(msg)
}
