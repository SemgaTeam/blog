package http

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/service"
	"github.com/SemgaTeam/blog/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"fmt"
)

type Server struct {
	echo *echo.Echo
	service Service
	conf *config.Config
}

type Service struct {
	post service.PostService
}

func NewEchoServer(conf *config.Config, db *gorm.DB) (Server, error) {
	echo := echo.New()

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)

	s := Server{
		echo,
		Service{
			postService,
		},
		conf,
	}

	s.setupRouter()

	return s, nil
}

func (s Server) setupRouter() {
	s.echo.Pre(middleware.RemoveTrailingSlash())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURIPath:    true,
		LogMethod: true,
		LogError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				fmt.Printf("Error %s: %v %v %v\n", v.Error.Error(), v.Method, v.URIPath, v.Status)
			}
			fmt.Printf("%v %v %v\n", v.Method, v.URIPath, v.Status)
			return nil
		},
	}))
	s.echo.HTTPErrorHandler = ErrorHandler

	api := s.echo.Group("/api")
	posts := api.Group("/post")

	posts.GET("/:id", s.GetPost)
	posts.POST("/", s.CreatePost)
	posts.PUT("/:id", s.UpdatePost)
	posts.DELETE("/:id", s.DeletePost)
}

func (s Server) Start() {
	s.echo.Logger.Fatal(
		s.echo.Start(
			fmt.Sprintf("%s:%s", 
				s.conf.App.Address, 
				s.conf.App.Port,
			),
		),
	)
}
