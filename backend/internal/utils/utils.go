package utils

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"context"
	"net/http"
	"strconv"
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

func HandleSorting(q *gorm.DB, s dto.Sorting, allowedFields []string) error { // handle sorting requests
	if s.SortField == "" { 
		return nil
	}

	allowed := false
	for _, allowedField := range allowedFields {
		if s.SortField == allowedField {
			allowed = true
		}
	}

	if !allowed {
		return e.BadRequest(nil, "sort field is not allowed")
	}

	q = q.Order(s.SortField + " " + s.SortOrder)

	return nil
}

func HandlePagination(q *gorm.DB, p dto.Pagination) {
	if p.Page != 0 && p.PerPage != 0 {
		q = q.
			Limit(p.PerPage).
			Offset(
				(p.Page - 1)*p.PerPage,
			)
	}
}

func GetClaimsFromContext(c echo.Context, tokenType string) (*entities.Claims, error) {
	token := c.Get(tokenType).(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims) 
	if ok != true {
		return nil, e.Unauthorized(errors.New("no claims"), "token is invalid")
	}

	return claims, nil
}
