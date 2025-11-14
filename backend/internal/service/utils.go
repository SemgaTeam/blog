package service

import (
	"go.uber.org/zap"

	"context"
)

func FromContext(ctx context.Context) *zap.Logger {
	if v := ctx.Value("logger"); v != nil {
		if logger, ok := v.(*zap.Logger); ok {
			return logger
		}
	}

	return zap.L()
}
