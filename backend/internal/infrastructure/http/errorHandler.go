package http

import (
	d "github.com/SemgaTeam/blog/internal/error"
	h "github.com/SemgaTeam/blog/internal/infrastructure/http/error"
	"github.com/labstack/echo/v4"

	"errors"
	"fmt"
)

func ErrorHandler(err error, c echo.Context) {
	var httpErr *h.HTTPError
	var echoHttpErr *echo.HTTPError

	switch {
	case errors.As(err, &httpErr):
		break

	case errors.As(err, &echoHttpErr):
		httpErr = &h.HTTPError{
			Code:    echoHttpErr.Code,
			Message: fmt.Sprint(echoHttpErr.Message),
			Err:     err,
		}

	case errors.Is(err, d.ErrPostNotFound):
		httpErr = h.BadRequest(err, "post not found")

	case errors.Is(err, d.ErrUserNotFound):
		httpErr = h.BadRequest(err, "user not found")

	case errors.Is(err, d.ErrInvalidCredentials):
		httpErr = h.Unauthorized(err, "invalid credentials")

	default:
		httpErr = h.Internal(err)
	}

	if !c.Response().Committed {
		c.JSON(httpErr.Code, map[string]string{
			"error": httpErr.Message,
		})
	}
}
