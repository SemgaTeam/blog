package service

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/log"
	"github.com/SemgaTeam/blog/internal/repository"
	"go.uber.org/zap"
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
	user, err := s.repo.user.CreateUser(name, password)
	if err != nil {
		log.Log.Info("create user error", zap.Error(err))
		return nil, err
	}
	
	log.Log.Debug("created user", zap.Int("id", user.ID))
	return user, nil
}

func (s *userService) GetUser(id int) (*entities.User, error) {
	user, err := s.repo.user.GetUser(id)
	if err != nil {
		log.Log.Info("get user error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}
	
	log.Log.Debug("got user", zap.Int("id", user.ID))
	return user, nil
}

func (s *userService) UpdateUser(id int, name, password string) (*entities.User, error) {
	user, err := s.repo.user.UpdateUser(id, name, password)
	if err != nil {
		log.Log.Info("update user error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}
	
	log.Log.Debug("updated user", zap.Int("id", user.ID))
	return user, nil
}

func (s *userService) DeleteUser(id int) (int, error) {
	_, err := s.repo.user.DeleteUser(id)
	if err != nil {
		log.Log.Info("delete user error", zap.Error(err), zap.Int("id", id))
		return 0, err
	}
	
	log.Log.Debug("deleted user", zap.Int("id", id))
	return id, nil
}
