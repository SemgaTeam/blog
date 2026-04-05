package main

import (
	"github.com/SemgaTeam/blog/internal/application"
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/domain/usecases"
	"github.com/SemgaTeam/blog/internal/infrastructure/db"
	"github.com/SemgaTeam/blog/internal/infrastructure/http"
	"github.com/SemgaTeam/blog/internal/infrastructure/repository"
	"github.com/SemgaTeam/blog/internal/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.GetConfig()
	log.InitLogger("logs/log")
	defer func(){ 
		_ = log.Log.Sync()
	}()

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

	postRepo := repository.NewPostRepository(db)
	log.Log.Debug("initialized post repository")

	postService := usecases.NewPostService(postRepo)
	log.Log.Debug("initialized post service")

	userRepo := repository.NewUserRepository(db)
	log.Log.Debug("initialized user repository")

	hashRepo := repository.NewHashRepository(conf.Hash)
	log.Log.Debug("initialized hash repository")

	userService := usecases.NewUserService(userRepo)
	log.Log.Debug("initialized user service")

	tokenRepo, err := repository.NewTokenRepository(conf)
	if err != nil {
		log.Log.Fatal("token repository initialization error", zap.Error(err))
		panic(err)
	}
	log.Log.Debug("initialized token repository")

	authService, err := usecases.NewAuthService(conf.Auth, tokenRepo, userRepo, hashRepo)
	if err != nil {
		log.Log.Fatal("auth service initialization error", zap.Error(err))
		panic(err)
	}
	log.Log.Debug("initialized auth service")

	service := application.NewService(postService, userService, authService)

	s, err := http.NewEchoServer(conf, service)
	if err != nil {
		log.Log.Fatal("failed to initialize echo server", zap.Error(err))
		panic(err)
	}
	log.Log.Info("initialized echo server")

	log.Log.Info("server started", zap.String("port", conf.App.Port), zap.String("address", conf.App.Address))
	s.Start()
}
