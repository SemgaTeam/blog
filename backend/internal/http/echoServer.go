package http

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/service"
	"github.com/SemgaTeam/blog/internal/log"
	"github.com/SemgaTeam/blog/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
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
	user service.UserService
}

func NewEchoServer(conf *config.Config, db *gorm.DB) (Server, error) {
	echo := echo.New()

	postRepo := repository.NewPostRepository(db)
	log.Log.Debug("initialized post repository")

	postService := service.NewPostService(postRepo)
	log.Log.Debug("initialized post service")

	userRepo := repository.NewUserRepository(db)
	log.Log.Debug("initialized user repository")

	userService := service.NewUserService(userRepo)
	log.Log.Debug("initialized user service")

	s := Server{
		echo,
		Service{
			postService,
			userService,
		},
		conf,
	}

	s.setupRouter()
	log.Log.Info("setup router completed")

	return s, nil
}

func (s Server) setupRouter() {
	s.echo.Pre(middleware.RemoveTrailingSlash())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURIPath: true,
		LogMethod: true,
		LogError: true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		if v.Error != nil {
			log.Log.Info(fmt.Sprintf("%v %v %v", v.Method, v.URIPath, v.Status), zap.Error(v.Error), zap.String("request_id", v.RequestID))
		}	else {
			log.Log.Info(fmt.Sprintf("%v %v %v", v.Method, v.URIPath, v.Status), zap.String("request_id", v.RequestID))
		}
		return nil
		},
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(SetLoggerMiddleware(log.Log))

	s.echo.HTTPErrorHandler = ErrorHandler

	api := s.echo.Group("/api")
	posts := api.Group("/post")
	users := api.Group("/user")

	posts.GET("/:id", s.GetPost)
	posts.GET("", s.GetPosts)
	posts.POST("", s.CreatePost)
	posts.PUT("/:id", s.UpdatePost)
	posts.DELETE("/:id", s.DeletePost)

	users.GET("/:id", s.GetUserById)
	users.POST("", s.CreateUser)
	users.PUT("/:id", s.UpdateUser)
	users.DELETE("/:id", s.DeleteUser)
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
