package service

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/utils"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/SemgaTeam/blog/internal/repository"

	"context"
)

type AuthService interface {
	LogIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error)
	SignIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error)
	RefreshTokens(context.Context, int) (*entities.AuthToken, *entities.AuthToken, error)
}

type authService struct {
	repo authServiceRepo	
	conf *config.Auth
}

type authServiceRepo struct {
	token repository.TokenRepository
	user repository.UserRepository
}

func NewAuthService(conf *config.Auth, tokenRepo repository.TokenRepository, userRepo repository.UserRepository) (AuthService, error) {
	return &authService{
		repo: authServiceRepo{
			token: tokenRepo,
			user: userRepo,
		},
		conf: conf,
	}, nil
}

func (s *authService) LogIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	user, err := s.repo.user.GetUserByName(name)	
	if err != nil {
		return nil, nil, err
	}

	if user.Password != password {
		return nil, nil, e.ErrInvalidCredentials
	}

	authToken, refreshToken, err := s.generateTokens(user.ID, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *authService) SignIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error) {
	user, err := s.repo.user.CreateUser(name, password)
	if err != nil {
		return nil, nil, err
	}

	authToken, refreshToken, err := s.generateTokens(user.ID, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *authService) RefreshTokens(ctx context.Context, userId int) (*entities.AuthToken, *entities.AuthToken, error) {
	authToken, refreshToken, err := s.generateTokens(userId, s.conf.AccessExpirationSecs, s.conf.RefreshExpirationSecs)
	if err != nil {
		return nil, nil, err
	}

	return authToken, refreshToken, nil
}

func (s *authService) generateTokens(userId int, accessExpirationSecs, refreshExpirationSecs int) (*entities.AuthToken, *entities.AuthToken, error) {
	authClaims := utils.GetClaims(userId, accessExpirationSecs)
	refreshClaims := utils.GetClaims(userId, refreshExpirationSecs)

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
