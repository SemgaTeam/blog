package http

import (
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/utils"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"net/http"
)

func (s Server) LogIn(c echo.Context) error {
	var request dto.LogInRequest	
	ctx := c.Request().Context()
	log := utils.FromContext(ctx)

	if err := c.Bind(&request); err != nil {
		return e.BadRequest(err, "invalid request body")
	}

	accessToken, refreshToken, err := s.service.auth.LogIn(ctx, request.Name, request.Password)
	if err != nil {
		log.Debug("log in error", zap.Error(err))
		return err
	}

	accessCookie := utils.SetAuthCookie("accessToken", accessToken.Value, accessToken.Claims.ExpiresAt.Time)
	refreshCookie := utils.SetAuthCookie("refreshToken", refreshToken.Value, refreshToken.Claims.ExpiresAt.Time)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.NoContent(http.StatusNoContent)
}

func (s Server) SignIn(c echo.Context) error {
	var request dto.SignInRequest	
	ctx := c.Request().Context()
	log := utils.FromContext(ctx)

	if err := c.Bind(&request); err != nil {
		return e.BadRequest(err, "invalid request body")
	}

	accessToken, refreshToken, err := s.service.auth.SignIn(ctx, request.Name, request.Password)
	if err != nil {
		log.Debug("sign in error", zap.Error(err))
		return err
	}

	accessCookie := utils.SetAuthCookie("accessToken", accessToken.Value, accessToken.Claims.ExpiresAt.Time)
	refreshCookie := utils.SetAuthCookie("refreshToken", refreshToken.Value, refreshToken.Claims.ExpiresAt.Time)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.NoContent(http.StatusNoContent)
}

func (s Server) LogOut(c echo.Context) error {
	var accessCookie, refreshCookie http.Cookie

	accessCookie.MaxAge = -1
	refreshCookie.MaxAge = -1

	c.SetCookie(&accessCookie)
	c.SetCookie(&refreshCookie)

	return c.NoContent(http.StatusNoContent)
}
