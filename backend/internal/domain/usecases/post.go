package usecases

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/domain"
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/utils"
	"go.uber.org/zap"

	"context"
)

type PostUseCase struct {
	repo domain.PostRepository
}

func NewPostUseCase(postRepo domain.PostRepository) *PostUseCase {
	return &PostUseCase{
		repo: postRepo,
	}
}

func (s *PostUseCase) CreatePost(ctx context.Context, name, contents string, authorId int) (*entities.Post, error) {
	log := utils.GetLoggerFromContext(ctx)

	post, err := entities.NewPost(name, contents, authorId)
	if err != nil {
		return nil, err
	}

	if err = s.repo.Save(post); err != nil {
		log.Info("create post error", zap.Error(err))
		return nil, err
	}

	log.Debug("created post", zap.Int("id", post.ID))
	return post, nil
}

func (s *PostUseCase) GetPost(ctx context.Context, id int) (*entities.Post, error) {
	log := utils.GetLoggerFromContext(ctx)

	post, err := s.repo.ById(id)
	if err != nil {
		log.Info("get post error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}

	log.Debug("got post", zap.Int("id", post.ID))
	return post, nil
}

func (s *PostUseCase) GetPosts(ctx context.Context, params dto.GetPostParams) ([]entities.Post, int64, error) {
	log := utils.GetLoggerFromContext(ctx)

	pagination, err := domain.NewPagination(params.Pagination.Page, params.Pagination.PerPage)
	if err != nil {
		return nil, 0, err
	}

	sorting, err := domain.NewPostSorting(params.Sorting.SortField, params.Sorting.SortOrder)
	if err != nil {
		return nil, 0, err
	}

	validatedParams := domain.GetPostParams{
		IDs:        params.IDs,
		Name:       params.Name,
		AuthorID:   params.AuthorID,
		Pagination: *pagination,
		Sorting:    *sorting,
	}

	posts, total, err := s.repo.ByParams(validatedParams)
	if err != nil {
		log.Info("get posts error", zap.Error(err))
		return nil, 0, err
	}

	log.Debug("got posts", zap.Int64("total", total))
	return posts, total, nil
}

func (s *PostUseCase) UpdatePost(ctx context.Context, id int, name, contents string) (*entities.Post, error) {
	log := utils.GetLoggerFromContext(ctx)

	post, err := entities.UpdatePost(id, name, contents)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(post); err != nil {
		log.Info("update post error", zap.Error(err))
		return nil, err
	}

	log.Debug("updated post", zap.Int("id", post.ID))
	return post, nil
}

func (s *PostUseCase) DeletePost(ctx context.Context, id int) (int, error) {
	log := utils.GetLoggerFromContext(ctx)

	if err := s.repo.Delete(id); err != nil {
		log.Info("delete post error", zap.Error(err))
		return 0, err
	}

	log.Debug("deleted post", zap.Int("id", id))
	return id, nil
}
