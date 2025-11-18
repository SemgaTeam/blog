package utils

import (
	"github.com/SemgaTeam/blog/internal/entities"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
	"errors"
)

func GetLoggerFromContext(ctx context.Context) *zap.Logger { // get logger from context
	if v := ctx.Value("logger"); v != nil {
		if logger, ok := v.(*zap.Logger); ok {
			return logger
		}
	}

	return zap.L()
}

func GetClaims(userId int, expirationSecs int) entities.Claims { // get jwt claims for app use
	date := jwt.NewNumericDate(
		time.Now().Add(
			time.Duration(expirationSecs)*time.Second,
		),
	)

	return entities.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.Itoa(userId),
			ExpiresAt: date,
		},
	}	
}

func SetAuthCookie(name, value string, expires time.Time) *http.Cookie { 
	var c http.Cookie

	c.Name = name
	c.Value = value
	c.Expires = expires
	c.HttpOnly = true
	c.Secure = true

	return &c
}

func HandleSorting(q *gorm.DB, sortField, sortOrder string, allowedFields []string) error { // handle sorting requests
	sortField = strings.TrimSpace(strings.ToLower(sortField))
	sortOrder = strings.TrimSpace(strings.ToLower(sortOrder))

	if sortField == "" { 
		return nil
	}

	if sortOrder == "" {
		sortOrder = "asc" // default value
	}

	allowed := false
	for _, allowedField := range allowedFields {
		if sortField == allowedField {
			allowed = true
		}
	}

	if !allowed {
		return e.BadRequest(nil, "sort field is not allowed")
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		return e.BadRequest(nil, "sort order is invalid")
	}

	q = q.Order(sortField + " " + sortOrder)

	return nil
}

func GetClaimsFromContext(c echo.Context, tokenType string) (*entities.Claims, error) {
	token := c.Get(tokenType).(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims) 
	if ok != true {
		return nil, e.Internal(errors.New("no claims"))
	}

	return claims, nil
}
