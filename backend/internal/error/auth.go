package error

import (
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)
