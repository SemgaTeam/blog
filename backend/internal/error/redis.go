package error

import (
	"errors"
)

var (
	ErrRedisTokenNotFound = errors.New("redis token not found")
	ErrRedisTokenSetFailed = errors.New("failed to set token")
	ErrRedisTokenDeleteFailed = errors.New("failed to delete token")
)
