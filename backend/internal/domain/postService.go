package domain

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/utils"
	"go.uber.org/zap"

	"context"
)

type PostServiceRepo struct {
	post PostRepository
}

type PostService struct {
	repo PostServiceRepo
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{
		repo: PostServiceRepo{
			postRepo,
		},
	}
}

func (s *PostService) CreatePost(ctx context.Context, name, contents string, authorId int) (*entities.Post, error) {
	log := utils.GetLoggerFromContext(ctx)

	post, err := entities.NewPost(name, contents, authorId)
	if err != nil {
		return nil, err
	}

	if err = s.repo.post.Save(post); err != nil {
		log.Info("create post error", zap.Error(err))
		return nil, err
	}

	log.Debug("created post", zap.Int("id", post.ID))
	return post, nil
}

func (s *PostService) GetPost(ctx context.Context, id int) (*entities.Post, error) {
	log := utils.GetLoggerFromContext(ctx)

	post, err := s.repo.post.ById(id)
	if err != nil {
		log.Info("get post error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}

	log.Debug("got post", zap.Int("id", post.ID))
	return post, nil
}

func (s *PostService) GetPosts(ctx context.Context, params dto.GetPostParams) ([]entities.Post, int64, error) {
	log := utils.GetLoggerFromContext(ctx)

	posts, total, err := s.repo.post.ByParams(params)
	if err != nil {
		log.Info("get posts error", zap.Error(err))
		return nil, 0, err
	}

	log.Debug("got posts", zap.Int64("total", total))
	return posts, total, nil
}

func (s *PostService) UpdatePost(ctx context.Context, id int, name, contents string) (*entities.Post, error) {
	log := utils.GetLoggerFromContext(ctx)

	post, err := entities.UpdatePost(id, name, contents)
	if err != nil {
		return nil, err
	}

	if err := s.repo.post.Save(post); err != nil {
		log.Info("update post error", zap.Error(err))
		return nil, err
	}

	log.Debug("updated post", zap.Int("id", post.ID))
	return post, nil
}

func (s *PostService) DeletePost(ctx context.Context, id int) (int, error) {
	log := utils.GetLoggerFromContext(ctx)

	if err := s.repo.post.Delete(id); err != nil {
		log.Info("delete post error", zap.Error(err))
		return 0, err
	}

	log.Debug("deleted post", zap.Int("id", id))
	return id, nil
}
