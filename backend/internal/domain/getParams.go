package domain

type GetUserParams struct {
	IDs        []int
	Name       string
	Pagination pagination
	Sorting    sorting
}

type GetPostParams struct {
	IDs        []int
	Name       string
	AuthorID   int
	Pagination pagination
	Sorting    sorting
}
