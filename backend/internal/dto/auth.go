package dto

type LogInRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
