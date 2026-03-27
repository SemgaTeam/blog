package domain

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/utils"
	"go.uber.org/zap"

	"context"
)

type UserService struct {
	user UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		user: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name, password string) (*entities.User, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := entities.NewUser(name, password)
	if err != nil {
		return nil, err
	}

	if err = s.user.Save(user); err != nil {
		log.Info("create user error", zap.Error(err))
		return nil, err
	}

	log.Debug("created user", zap.Int("id", user.ID))
	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, id int) (*entities.User, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := s.user.ById(id)
	if err != nil {
		log.Info("get user error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}

	log.Debug("got user", zap.Int("id", user.ID))
	return user, nil
}

func (s *UserService) GetUsers(ctx context.Context, params dto.GetUserParams) ([]entities.User, int64, error) {
	log := utils.GetLoggerFromContext(ctx)

	pagination, err := NewPagination(params.Pagination.Page, params.Pagination.PerPage)
	if err != nil {
		return nil, 0, err
	}

	sorting, err := NewPostSorting(params.Sorting.SortField, params.Sorting.SortOrder)
	if err != nil {
		return nil, 0, err
	}

	validatedParams := GetUserParams{
		IDs:        params.IDs,
		Name:       params.Name,
		Pagination: *pagination,
		Sorting:    *sorting,
	}

	users, total, err := s.user.ByParams(validatedParams)
	if err != nil {
		log.Info("get posts error", zap.Error(err))
		return nil, 0, err
	}

	log.Debug("got users", zap.Int64("total", total))
	return users, total, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, name, password string) (*entities.User, error) {
	log := utils.GetLoggerFromContext(ctx)

	user, err := entities.UpdateUser(id, name, password)
	if err != nil {
		return nil, err
	}

	if err := s.user.Save(user); err != nil {
		log.Info("update user error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}

	log.Debug("updated user", zap.Int("id", user.ID))
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) (int, error) {
	log := utils.GetLoggerFromContext(ctx)

	if err := s.user.Delete(id); err != nil {
		log.Info("delete user error", zap.Error(err), zap.Int("id", id))
		return 0, err
	}

	log.Debug("deleted user", zap.Int("id", id))
	return id, nil
}
