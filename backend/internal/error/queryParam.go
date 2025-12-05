package error

import "errors"

var (
	ErrInvalidQueryParam = errors.New("invalid query param")
)
