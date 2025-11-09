package http

import (
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/labstack/echo/v4"

	"errors"
)

func ErrorHandler(err error, c echo.Context) {
	var appErr *e.AppError

	switch {
	case errors.Is(err, e.ErrPostNotFound):
		appErr = e.BadRequest(err, "post not found")
	default:
		appErr = e.Internal(err)	
	}

	if !c.Response().Committed {
		c.JSON(appErr.Code, map[string]string{
			"error": appErr.Message,
		})
	}
}
