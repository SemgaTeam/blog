package error

var (
	ErrInvalidSortingField = NewError("invalid sorting field")
	ErrInvalidSortOrder    = NewError("invalid sort order")
)
