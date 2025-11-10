package dto

type Post struct {
	ID int `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name string `json:"name"`
	Contents string `json:"contents"`
	AuthorID int `json:"author_id"`
}

type CreatePostRequest struct {
	Name string `json:"name"`	
	Contents string `json:"contents"`
	AuthorID int `json:"author_id"`
}

type CreatePostResponse = Post

type GetPostResponse = Post

type UpdatePostRequest struct {
	Name string `json:"name"`
	Contents string `json:"contents"`
}

type UpdatePostResponse = Post

type DeletePostResponse struct {
	ID int `json:"id"`
}

type GetPostParams struct {
	IDs []int `query:"ids"`
	Page int `query:"page"`
	PerPage int `query:"perPage"`
	SortField string `query:"sortField"`
	SortOrder string `query:"sortOrder"`
	Name string `query:"name"`
	AuthorID int `query:"authorId"`
}

type GetPostsResponse struct {
	Data []Post `json:"data"`
	Total int64 `json:"total"`
}
