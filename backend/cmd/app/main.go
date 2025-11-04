package main

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/http"
)

func main() {
	conf := config.GetConfig()	

	s, err := http.NewEchoServer(conf)
	if err != nil {
		panic(err)
	}

	s.Start()
}
