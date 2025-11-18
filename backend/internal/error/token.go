package error

import (
	"errors"
)

var (
	ErrTokenSigningMethodNotAllowed = errors.New("signing method not allowed")
	ErrSigningToken = errors.New("error signing token")
)
