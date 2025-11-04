package http

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/labstack/echo/v4"

	"fmt"
)

type Server struct {
	echo *echo.Echo
	service Service
	conf *config.Config
}

type Service struct {}

func NewEchoServer(conf *config.Config) (Server, error) {
	echo := echo.New()

	return Server{
		echo,
		Service{},
		conf,
	}, nil
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
