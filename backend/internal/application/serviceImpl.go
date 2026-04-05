package application

import (
	"github.com/SemgaTeam/blog/internal/domain/usecases"
)

type service struct {
	*usecases.AuthUseCase
	*usecases.UserUseCase
	*usecases.PostUseCase
}

func NewService(postUseCase *usecases.PostUseCase, userUseCase *usecases.UserUseCase, authUseCase *usecases.AuthUseCase) Service {
	return &service{
		AuthUseCase: authUseCase,
		UserUseCase: userUseCase,
		PostUseCase: postUseCase,
	}
}
