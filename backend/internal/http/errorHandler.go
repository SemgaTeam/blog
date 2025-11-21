package http

import (
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/labstack/echo/v4"

	"errors"
	"fmt"
)

func ErrorHandler(err error, c echo.Context) {
	var appErr *e.AppError
	var httpErr *echo.HTTPError

	switch {
	case errors.As(err, &appErr):
		break
	
	case errors.As(err, &httpErr):
		appErr = &e.AppError{
			Code: httpErr.Code,
			Message: fmt.Sprint(httpErr.Message),
			Err: err,
		}

	case errors.Is(err, e.ErrPostNotFound):
		appErr = e.BadRequest(err, "post not found")

	case errors.Is(err, e.ErrUserNotFound):
		appErr = e.BadRequest(err, "user not found")

	case errors.Is(err, e.ErrInvalidCredentials):
		appErr = e.Unauthorized(err, "invalid credentials")

	case errors.Is(err, e.ErrUserAlreadyExists):
		appErr = e.BadRequest(err, "user already exists")

	default:
		appErr = e.Internal(err)
	}

	if !c.Response().Committed {
		c.JSON(appErr.Code, map[string]string{
			"error": appErr.Message,
		})
	}
}
