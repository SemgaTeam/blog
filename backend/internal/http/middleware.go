package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"context"
)


func SetLoggerMiddleware(baseLogger *zap.Logger) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)

			logger := baseLogger.With(zap.String("request_id", reqID))

			ctx := context.WithValue(c.Request().Context(), "requestId", reqID)
			ctx = context.WithValue(ctx, "logger", logger)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
