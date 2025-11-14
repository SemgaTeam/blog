package service

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/repository"
	"go.uber.org/zap"

	"context"
)

type UserService interface {
	CreateUser(context.Context, string, string) (*entities.User, error)
	GetUser(context.Context, int) (*entities.User, error)
	UpdateUser(context.Context, int, string, string) (*entities.User, error)
	DeleteUser(context.Context, int) (int, error)
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

func (s *userService) CreateUser(ctx context.Context, name, password string) (*entities.User, error) {
	log := FromContext(ctx)

	user, err := s.repo.user.CreateUser(name, password)
	if err != nil {
		log.Info("create user error", zap.Error(err))
		return nil, err
	}
	
	log.Debug("created user", zap.Int("id", user.ID))
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id int) (*entities.User, error) {
	log := FromContext(ctx)

	user, err := s.repo.user.GetUser(id)
	if err != nil {
		log.Info("get user error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}
	
	log.Debug("got user", zap.Int("id", user.ID))
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, name, password string) (*entities.User, error) {
	log := FromContext(ctx)

	user, err := s.repo.user.UpdateUser(id, name, password)
	if err != nil {
		log.Info("update user error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}
	
	log.Debug("updated user", zap.Int("id", user.ID))
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) (int, error) {
	log := FromContext(ctx)

	_, err := s.repo.user.DeleteUser(id)
	if err != nil {
		log.Info("delete user error", zap.Error(err), zap.Int("id", id))
		return 0, err
	}
	
	log.Debug("deleted user", zap.Int("id", id))
	return id, nil
}
