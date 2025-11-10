package error

import (
	"errors"
)

var (
	ErrPostNotFound = errors.New("post not found")
	ErrPostInvalidRequest = errors.New("post invalid request")
)
