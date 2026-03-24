package application

import (
	"github.com/SemgaTeam/blog/internal/domain"
)

type service struct {
	*domain.AuthService
	*domain.UserService
	*domain.PostService
}

func NewService(postService *domain.PostService, userService *domain.UserService, authService *domain.AuthService) Service {
	return &service{
		AuthService: authService,
		UserService: userService,
		PostService: postService,
	}
}
