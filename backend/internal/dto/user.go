package dto

type User struct {
	ID int `json:"id"`
	CreatedAt string `json:"created_at"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`	
	Password string `json:"password"`
}

type CreateUserResponse = User

type GetUserResponse = User

type GetUserParams struct {
	IDs []int `query:"ids"`
	Pagination Pagination `query:"pagination"`
	Sorting Sorting `query:"sorting"`
	Name string `query:"name"`
}

type GetUsersResponse struct {
	Data []User `json:"data"`
	Total int64 `json:"total"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserResponse = User

type DeleteUserResponse struct {
	ID int `json:"id"`
}
