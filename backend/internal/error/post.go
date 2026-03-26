package error

var (
	ErrPostNotFound       = NewError("post not found")
	ErrPostInvalidRequest = NewError("post invalid request")
)
