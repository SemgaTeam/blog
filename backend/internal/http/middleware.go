package http

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/golang-jwt/jwt/v5"

	"context"
	"net/http"
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

func GetAccessMiddleware(signingKey string, signingMethod string) echo.MiddlewareFunc {
	accessMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(signingKey),
		TokenLookup: "cookie:accessToken",
		ContextKey: "access",
		SigningMethod: signingMethod,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.Claims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
		},
	})

	return accessMiddleware
}

func GetRefreshMiddleware(signingKey string, signingMethod string) echo.MiddlewareFunc {
	refreshMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(signingKey),
		TokenLookup: "cookie:refreshToken",
		ContextKey: "refresh",
		SigningMethod: signingMethod,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.Claims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
		},
	})

	return refreshMiddleware
}
