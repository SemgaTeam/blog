package error

var (
	ErrUserAlreadyExists = NewError("user already exists")
	ErrUserNotFound      = NewError("user not found")
	ErrInvalidUser       = NewError("invalid user")
)
