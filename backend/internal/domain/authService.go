package domain

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/entities"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/infrastructure/repository"
	"github.com/SemgaTeam/blog/internal/utils"
	"go.uber.org/zap"

	"context"
)

type AuthService interface {
	LogIn(context.Context, string, string) (*entities.AuthToken, *entities.AuthToken, error)
	SignIn(context.Context, string, string) (*entities.AuthToken, *entities.AuthToken, error)
	RefreshTokens(context.Context, int, bool) (*entities.AuthToken, *entities.AuthToken, error)
	GetMe(context.Context, int) (*entities.User, error)
}

type authService struct {
	repo authServiceRepo
	conf *config.Auth
}

type authServiceRepo struct {
	token repository.TokenRepository
	user  repository.UserRepository
	hash  repository.HashRepository
}

func NewAuthService(conf *config.Auth, tokenRepo repository.TokenRepository, userRepo repository.UserRepository, hashRepo repository.HashRepository) (AuthService, error) {
	return &authService{
		repo: authServiceRepo{
			token: tokenRepo,
			user:  userRepo,
			hash:  hashRepo,
		},
		conf: conf,
	}, nil
}

func (s *authService) LogIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.repo.user.GetUserByName(name)
	if err != nil {
		log.Info("user not found by name", zap.Error(err))
		return nil, nil, err
	}

	if !s.repo.hash.IsPasswordValid(password, user.Password) {
		log.Info("invalid password", zap.Int("user_id", user.ID))
		return nil, nil, e.ErrInvalidCredentials
	}

	authToken, refreshToken, err := s.generateTokens(user.ID, user.IsAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		log.Info("error generating tokens", zap.Error(err))
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *authService) SignIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	hashedPassword, err := s.repo.hash.HashPassword(password)
	if err != nil {
		log.Error("hashing failed", zap.Error(err))
		return nil, nil, err
	}

	user, err := s.repo.user.CreateUser(name, hashedPassword)
	if err != nil {
		return nil, nil, err
	}

	authToken, refreshToken, err := s.generateTokens(user.ID, user.IsAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *authService) RefreshTokens(ctx context.Context, userId int, isAdmin bool) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	authToken, refreshToken, err := s.generateTokens(userId, isAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		log.Info("Error generating tokens", zap.Error(err))
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *authService) GetMe(ctx context.Context, userId int) (*entities.User, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.repo.user.GetUserById(userId)
	if err != nil {
		log.Info("failed getting user", zap.Error(err), zap.Int("id", userId))
		return nil, err
	}
	log.Debug("got user info", zap.Int("id", userId))

	return user, nil
}

func (s *authService) generateTokens(userId int, isAdmin bool, accessExpirationSecs, refreshExpirationSecs int) (*entities.AuthToken, *entities.AuthToken, error) {
	authClaims := utils.GetClaims(userId, isAdmin, accessExpirationSecs)
	refreshClaims := utils.GetClaims(userId, isAdmin, refreshExpirationSecs)

	authToken, err := s.repo.token.GenerateAndSignToken(authClaims)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.repo.token.GenerateAndSignToken(refreshClaims)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}
