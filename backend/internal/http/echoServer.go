package http

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/log"
	"github.com/SemgaTeam/blog/internal/repository"
	"github.com/SemgaTeam/blog/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
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
	auth service.AuthService
}

func NewEchoServer(conf *config.Config, db *gorm.DB, rdb *redis.Client) (*Server, error) {
	echo := echo.New()

	postRepo := repository.NewPostRepository(db)
	log.Log.Debug("initialized post repository")

	postService := service.NewPostService(postRepo)
	log.Log.Debug("initialized post service")

	userRepo := repository.NewUserRepository(db)
	log.Log.Debug("initialized user repository")

	hashRepo := repository.NewHashRepository(conf.Hash)
	log.Log.Debug("initialized hash repository")

	userService := service.NewUserService(userRepo)
	log.Log.Debug("initialized user service")

	tokenRepo, err := repository.NewTokenRepository(conf)
	if err != nil {
		log.Log.Fatal("token repository initialization error", zap.Error(err))
		return nil, err
	}
	log.Log.Debug("initialized token repository")

	redisRepo := repository.NewRedisRepository(rdb)
	log.Log.Debug("initialized redis repository")

	authService, err := service.NewAuthService(conf.Auth, tokenRepo, userRepo, hashRepo, redisRepo)
	if err != nil {
		log.Log.Fatal("auth service initialization error", zap.Error(err))
		return nil, err
	}
	log.Log.Debug("initialized auth service")

	s := Server{
		echo,
		Service{
			postService,
			userService,
			authService,
		},
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
	users.POST("", s.CreateUser)

	usersAuth.PUT("/:id", s.UpdateUser)
	usersAuth.DELETE("/:id", s.DeleteUser)

	auth.POST("/signin", s.SignIn)
	auth.POST("/login", s.LogIn)
	auth.POST("/logout", s.LogOut, accessMiddleware)
	auth.POST("/refresh", s.RefreshTokens, refreshMiddleware)
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
