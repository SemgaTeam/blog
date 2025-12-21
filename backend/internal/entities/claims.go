package entities

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	IsAdmin bool
	jwt.RegisteredClaims
}
