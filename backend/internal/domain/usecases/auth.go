package usecases

import (
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/domain"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/utils"
	"go.uber.org/zap"

	"context"
)

type AuthUseCase struct {
	conf *config.Auth
	token domain.TokenRepository
	user  domain.UserRepository
	hash  domain.HashRepository
}

func NewAuthUseCase(conf *config.Auth, tokenRepo domain.TokenRepository, userRepo domain.UserRepository, hashRepo domain.HashRepository) (*AuthUseCase, error) {
	return &AuthUseCase{
		conf: conf,
		token: tokenRepo,
		user:  userRepo,
		hash:  hashRepo,
	}, nil
}

func (s *AuthUseCase) LogIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.user.ByName(name)
	if err != nil {
		log.Info("user not found by name", zap.Error(err))
		return nil, nil, err
	}

	if !s.hash.IsPasswordValid(password, user.Password) {
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

func (s *AuthUseCase) SignIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	hashedPassword, err := s.hash.HashPassword(password)
	if err != nil {
		log.Error("hashing failed", zap.Error(err))
		return nil, nil, err
	}

	user, err := entities.NewUser(name, hashedPassword)
	if err != nil {
		return nil, nil, err
	}

	if err = s.user.Save(user); err != nil {
		return nil, nil, err
	}

	authToken, refreshToken, err := s.generateTokens(user.ID, user.IsAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *AuthUseCase) RefreshTokens(ctx context.Context, userId int, isAdmin bool) (*entities.AuthToken, *entities.AuthToken, error) {
	log := utils.GetLoggerFromContext(ctx)

	authToken, refreshToken, err := s.generateTokens(userId, isAdmin, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		log.Info("Error generating tokens", zap.Error(err))
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *AuthUseCase) GetMe(ctx context.Context, userId int) (*entities.User, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.user.ById(userId)
	if err != nil {
		log.Info("failed getting user", zap.Error(err), zap.Int("id", userId))
		return nil, err
	}
	log.Debug("got user info", zap.Int("id", userId))

	return user, nil
}

func (s *AuthUseCase) generateTokens(userId int, isAdmin bool, accessExpirationSecs, refreshExpirationSecs int) (*entities.AuthToken, *entities.AuthToken, error) {
	authClaims := utils.GetClaims(userId, isAdmin, accessExpirationSecs)
	refreshClaims := utils.GetClaims(userId, isAdmin, refreshExpirationSecs)

	authToken, err := s.token.GenerateAndSignToken(authClaims)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.token.GenerateAndSignToken(refreshClaims)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}
