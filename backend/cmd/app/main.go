package main

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/http"
	"github.com/SemgaTeam/blog/internal/db"
)

func main() {
	conf := config.GetConfig()	

	if err := db.RunMigrations(conf.Postgres, "migrations"); err != nil {
		panic(err)
	}

	db, err := db.NewPostgresConnection(conf.Postgres)
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	s, err := http.NewEchoServer(conf, db)
	if err != nil {
		panic(err)
	}

	s.Start()
}
