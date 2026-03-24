package http

import (
	"github.com/SemgaTeam/blog/internal/application"
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"fmt"
)

type Server struct {
	echo    *echo.Echo
	service application.Service
	conf    *config.Config
}

func NewEchoServer(conf *config.Config, service application.Service) (*Server, error) {
	echo := echo.New()

	s := Server{
		echo,
		service,
		conf,
	}

	s.setupRouter()
	log.Log.Info("setup router completed")

	return &s, nil
}

func (s Server) setupRouter() {
	s.echo.Pre(middleware.RemoveTrailingSlash())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURIPath:   true,
		LogMethod:    true,
		LogError:     true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Log.Info(fmt.Sprintf("%v %v %v", v.Method, v.URIPath, v.Status), zap.Error(v.Error), zap.String("request_id", v.RequestID))
			} else {
				log.Log.Info(fmt.Sprintf("%v %v %v", v.Method, v.URIPath, v.Status), zap.String("request_id", v.RequestID))
			}
			return nil
		},
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(SetLoggerMiddleware(log.Log))

	accessMiddleware := GetAccessMiddleware(s.conf.Auth.Secret, s.conf.Auth.SigningMethod)
	refreshMiddleware := GetRefreshMiddleware(s.conf.Auth.Secret, s.conf.Auth.SigningMethod)

	s.echo.HTTPErrorHandler = ErrorHandler

	api := s.echo.Group("/api")

	posts := api.Group("/post")
	postsAuth := posts.Group("", accessMiddleware)

	users := api.Group("/user")
	usersAuth := users.Group("", accessMiddleware)

	auth := api.Group("/auth")

	posts.GET("/:id", s.GetPost)
	posts.GET("", s.GetPosts)

	postsAuth.POST("", s.CreatePost)
	postsAuth.PUT("/:id", s.UpdatePost)
	postsAuth.DELETE("/:id", s.DeletePost)

	users.GET("/:id", s.GetUserById)
	users.GET("", s.GetUsers)
	users.POST("", s.CreateUser)

	usersAuth.PUT("/:id", s.UpdateUser)
	usersAuth.DELETE("/:id", s.DeleteUser)

	auth.POST("/signin", s.SignIn)
	auth.POST("/login", s.LogIn)
	auth.POST("/logout", s.LogOut)
	auth.POST("/refresh", s.RefreshTokens, refreshMiddleware)
	auth.POST("/me", s.GetMe, accessMiddleware)
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
