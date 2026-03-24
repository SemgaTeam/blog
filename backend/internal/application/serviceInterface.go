package application

type Service interface {
	AuthService
	PostService
	UserService
}

type AuthService interface {
}

type PostService interface {
}

type UserService interface {
}
