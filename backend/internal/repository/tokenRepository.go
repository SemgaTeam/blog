package repository

import (
	"github.com/SemgaTeam/blog/internal/entities"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenRepository interface {
	GenerateAndSignToken(entities.Claims) (*entities.AuthToken, error) 
}

type tokenRepository struct {
	conf *config.Auth
	signingMethod jwt.SigningMethod
}

func NewTokenRepository(conf *config.Config) (TokenRepository, error) {
	var signingMethod jwt.SigningMethod

	switch conf.Auth.SigningMethod {
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
	default:
		return nil, e.ErrTokenSigningMethodNotAllowed
	}

	return &tokenRepository{
		conf: conf.Auth,
		signingMethod: signingMethod,
	}, nil
}

func (r *tokenRepository) GenerateAndSignToken(claims entities.Claims) (*entities.AuthToken, error) {
	rawToken := jwt.NewWithClaims(r.signingMethod, claims)
	signedToken, err := r.signToken(rawToken)
	if err != nil {
		return nil, err
	}

	return &entities.AuthToken{
		Value: signedToken,
		Claims: claims,
	}, nil
}

func (r *tokenRepository) signToken(token *jwt.Token) (string, error) {
	tokenStr, err := token.SignedString([]byte(r.conf.Secret))
	if err != nil {
		return "", e.ErrSigningToken
	}

	return tokenStr, nil
}
