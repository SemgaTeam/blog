package http

import (
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
)

func (s Server) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var request dto.CreateUserRequest	

	if err := c.Bind(&request); err != nil {
		return e.BadRequest(err, "invalid request body")
	}

	user, err := s.service.user.CreateUser(ctx, request.Name, request.Password)
	if err != nil {
		return err
	}

	response := user.ToDTO()

	return c.JSON(http.StatusCreated, response)
}

func (s Server) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	var response dto.GetUserResponse
	
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.BadRequest(err, "invalid id")
	}

	user, err := s.service.user.GetUser(ctx, id)
	if err != nil {
		return err
	}

	response = user.ToDTO()

	return c.JSON(http.StatusOK, response)
}

func (s Server) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var response dto.UpdateUserResponse

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.BadRequest(err, "invalid id")
	}

	var request dto.UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return e.BadRequest(err, "invalid request")
	}

	user, err := s.service.user.UpdateUser(ctx, id, request.Name, request.Password)
	if err != nil {
		return err
	}

	response = user.ToDTO()

	return c.JSON(http.StatusOK, response)
}

func (s Server) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	var response dto.DeleteUserResponse

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.BadRequest(err, "invalid id")
	}

	_, err = s.service.user.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	response.ID = id

	return c.JSON(http.StatusOK, response)

}
