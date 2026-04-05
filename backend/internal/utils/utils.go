package utils

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"context"
	"net/http"
	"strconv"
	"time"
)

func GetLoggerFromContext(ctx context.Context) *zap.Logger { // get logger from context
	if v := ctx.Value("logger"); v != nil {
		if logger, ok := v.(*zap.Logger); ok {
			return logger
		}
	}

	return zap.L()
}

func GetClaims(userId int, isAdmin bool, expirationSecs int) entities.Claims { // get jwt claims for app use
	date := jwt.NewNumericDate(
		time.Now().Add(
			time.Duration(expirationSecs) * time.Second,
		),
	)

	return entities.Claims{
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(userId),
			ExpiresAt: date,
		},
	}
}

func SetAuthCookie(name, value, path string, expires time.Time) *http.Cookie {
	var c http.Cookie

	c.Name = name
	c.Value = value
	c.Expires = expires
	c.HttpOnly = true
	c.SameSite = http.SameSiteStrictMode
	c.Secure = true
	c.Path = path

	return &c
}

func HandlePagination(q *gorm.DB, page, perPage int) *gorm.DB {
	if page != 0 && perPage != 0 {
		q = q.
			Limit(perPage).
			Offset(
				(page - 1) * perPage,
			)
	}

	return q
}

func GetClaimsFromContext(c echo.Context, tokenType string) (*entities.Claims, error) {
	token := c.Get(tokenType).(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)
	if !ok {
		return nil, e.ErrUnauthorized
	}

	return claims, nil
}
