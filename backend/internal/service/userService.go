package service

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/repository"
)

type UserService interface {
	CreateUser(string, string) (*entities.User, error)
	GetUser(int) (*entities.User, error)
	UpdateUser(int, string, string) (*entities.User, error)
	DeleteUser(int) (int, error)
}

type userService struct {
	repo userServiceRepo
}

type userServiceRepo struct {
	user repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userServiceRepo{
			user: userRepo,
		},
	}
}

func (s *userService) CreateUser(name, password string) (*entities.User, error) {
	return s.repo.user.CreateUser(name, password)
}

func (s *userService) GetUser(id int) (*entities.User, error) {
	return s.repo.user.GetUser(id)
}

func (s *userService) UpdateUser(id int, name, password string) (*entities.User, error) {
	return s.repo.user.UpdateUser(id, name, password)
}

func (s *userService) DeleteUser(id int) (int, error) {
	return s.repo.user.DeleteUser(id)
}
