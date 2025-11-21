package main

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/db"
	"github.com/SemgaTeam/blog/internal/http"
	"github.com/SemgaTeam/blog/internal/log"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"fmt"
	"context"
)

func main() {
	conf := config.GetConfig()	
	log.InitLogger("logs/log")
	defer log.Log.Sync()

	migrationPath := "migrations"
	if err := db.RunMigrations(conf.Postgres, migrationPath); err != nil {
		log.Log.Fatal("failed to run migrations", zap.Error(err), zap.String("migration_path", migrationPath))
		panic(err)
	}
	log.Log.Info("migrations done", zap.String("migration_path", migrationPath))

	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Log.Fatal("failed to create postgresql connection", zap.Error(err))
		panic(err)
	}
	log.Log.Info("initialized PostgresQL connection")

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB: 0,
	})

	defer rdb.FlushDB(context.Background())

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Log.Fatal("failed to connect to redis", zap.Error(err), zap.String("password", conf.Redis.Password))
		panic(err)
	}

	s, err := http.NewEchoServer(conf, db, rdb)
	if err != nil {
		log.Log.Fatal("failed to initialize echo server", zap.Error(err))
		panic(err)
	}
	log.Log.Info("initialized echo server")

	log.Log.Info("server started", zap.String("port", conf.App.Port), zap.String("address", conf.App.Address))
	s.Start()
}
