package domain

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/domain/entities"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/utils"
	"go.uber.org/zap"

	"context"
)

type AuthService struct {
	repo AuthServiceRepo
	conf *config.Auth
}

type AuthServiceRepo struct {
	token TokenRepository
	user  UserRepository
	hash  HashRepository
}

func NewAuthService(conf *config.Auth, tokenRepo TokenRepository, userRepo UserRepository, hashRepo HashRepository) (*AuthService, error) {
	return &AuthService{
		repo: AuthServiceRepo{
			token: tokenRepo,
			user:  userRepo,
			hash:  hashRepo,
		},
		conf: conf,
	}, nil
}

func (s *AuthService) LogIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.repo.user.ByName(name)
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

func (s *AuthService) SignIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	hashedPassword, err := s.repo.hash.HashPassword(password)
	if err != nil {
		log.Error("hashing failed", zap.Error(err))
		return nil, nil, err
	}

	user, err := entities.NewUser(name, hashedPassword)
	if err != nil {
		return nil, nil, err
	}

	if err = s.repo.user.Save(user); err != nil {
		return nil, nil, err
	}

	authToken, refreshToken, err := s.generateTokens(user.ID, user.IsAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *AuthService) RefreshTokens(ctx context.Context, userId int, isAdmin bool) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	authToken, refreshToken, err := s.generateTokens(userId, isAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		log.Info("Error generating tokens", zap.Error(err))
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *AuthService) GetMe(ctx context.Context, userId int) (*entities.User, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.repo.user.ById(userId)
	if err != nil {
		log.Info("failed getting user", zap.Error(err), zap.Int("id", userId))
		return nil, err
	}
	log.Debug("got user info", zap.Int("id", userId))

	return user, nil
}

func (s *AuthService) generateTokens(userId int, isAdmin bool, accessExpirationSecs, refreshExpirationSecs int) (*entities.AuthToken, *entities.AuthToken, error) {
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
