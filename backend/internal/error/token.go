package error

var (
	ErrTokenSigningMethodNotAllowed = NewError("signing method not allowed")
	ErrSigningToken                 = NewError("error signing token")
)
