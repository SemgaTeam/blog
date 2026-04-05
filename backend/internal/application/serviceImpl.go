package application

import (
	"github.com/SemgaTeam/blog/internal/domain/usecases"
)

type service struct {
	*usecases.AuthService
	*usecases.UserService
	*usecases.PostService
}

func NewService(postService *usecases.PostService, userService *usecases.UserService, authService *usecases.AuthService) Service {
	return &service{
		AuthService: authService,
		UserService: userService,
		PostService: postService,
	}
}
