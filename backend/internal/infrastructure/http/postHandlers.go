package http

import (
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/infrastructure/http/error"
	"github.com/SemgaTeam/blog/internal/utils"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
)

func (s Server) CreatePost(c echo.Context) error {
	ctx := c.Request().Context()
	var request dto.CreatePostRequest

	claims, err := utils.GetClaimsFromContext(c, "access")
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return e.Unauthorized(err, "invalid user id")
	}

	var response dto.CreatePostResponse
	if err := c.Bind(&request); err != nil {
		return e.BadRequest(err, "invalid request")
	}

	post, err := s.service.CreatePost(ctx, request.Name, request.Contents, id)
	if err != nil {
		return err
	}

	response = post.ToDTO()

	return c.JSON(http.StatusCreated, response)
}

func (s Server) GetPost(c echo.Context) error {
	ctx := c.Request().Context()
	var response dto.GetPostResponse

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.BadRequest(err, "invalid id")
	}

	post, err := s.service.GetPost(ctx, id)
	if err != nil {
		return err
	}

	response = post.ToDTO()

	return c.JSON(http.StatusOK, response)
}

func (s Server) GetPosts(c echo.Context) error {
	ctx := c.Request().Context()
	var params dto.GetPostParams

	if err := c.Bind(&params); err != nil {
		return e.BadRequest(err, "invalid query parameters")
	}

	posts, total, err := s.service.GetPosts(ctx, params)
	if err != nil {
		return err
	}

	var response dto.GetPostsResponse

	for _, post := range posts {
		response.Data = append(response.Data, post.ToDTO())
	}

	response.Total = total

	return c.JSON(http.StatusOK, response)
}

func (s Server) UpdatePost(c echo.Context) error {
	ctx := c.Request().Context()
	var response dto.UpdatePostResponse

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.BadRequest(err, "invalid id")
	}

	var request dto.UpdatePostRequest
	if err := c.Bind(&request); err != nil {
		return e.BadRequest(err, "invalid request")
	}

	post, err := s.service.UpdatePost(ctx, id, request.Name, request.Contents)
	if err != nil {
		return err
	}

	response = post.ToDTO()

	return c.JSON(http.StatusOK, response)
}

func (s Server) DeletePost(c echo.Context) error {
	ctx := c.Request().Context()
	var response dto.DeletePostResponse

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.BadRequest(err, "invalid id")
	}

	_, err = s.service.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	response.ID = id

	return c.JSON(http.StatusOK, response)
}
