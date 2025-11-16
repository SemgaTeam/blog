package service

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"context"
	"strconv"
	"time"
)

func FromContext(ctx context.Context) *zap.Logger {
	if v := ctx.Value("logger"); v != nil {
		if logger, ok := v.(*zap.Logger); ok {
			return logger
		}
	}

	return zap.L()
}

func getClaims(userId int, expirationSecs int) entities.Claims {
	date := jwt.NewNumericDate(
		time.Now().Add(
			time.Duration(expirationSecs)*time.Second,
		),
	)

	return entities.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.Itoa(userId),
			ExpiresAt: date,
		},
	}	
}

func (s *authService) generateTokens(userId int, accessExpirationSecs, refreshExpirationSecs int) (*entities.AuthToken, *entities.AuthToken, error) {
	authClaims := getClaims(userId, accessExpirationSecs)
	refreshClaims := getClaims(userId, refreshExpirationSecs)

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
