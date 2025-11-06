package main

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/http"
	"github.com/SemgaTeam/blog/internal/repository"
)

func main() {
	conf := config.GetConfig()	

	db, err := repository.NewPostgresConnection(conf.Postgres)

	if err != nil {
		panic(err)
	}

	s, err := http.NewEchoServer(conf, db)
	if err != nil {
		panic(err)
	}

	s.Start()
}
