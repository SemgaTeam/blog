package application

type service struct {
	post PostService
	user UserService
	auth AuthService
}

func NewService(postService PostService, userService UserService, authService AuthService) Service {
	return &service{
		post: postService,
		user: userService,
		auth: authService,
	}
}
