package service 

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/repository"
	"go.uber.org/zap"

	"context"
)

type PostService interface {
	CreatePost(context.Context, string, string, int) (*entities.Post, error)
	GetPost(context.Context, int) (*entities.Post, error)
	GetPosts(context.Context, dto.GetPostParams) ([]entities.Post, int64, error)
	UpdatePost(context.Context, int, string, string) (*entities.Post, error)
	DeletePost(context.Context, int) (int, error)
}

type postServiceRepo struct {
	post repository.PostRepository
}

type postService struct {
	repo postServiceRepo
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		repo: postServiceRepo{
			postRepo,
		},
	}
}

func (s *postService) CreatePost(ctx context.Context, name, contents string, authorId int) (*entities.Post, error) {
	log := FromContext(ctx)

	post, err := s.repo.post.CreatePost(name, contents, authorId)
	if err != nil {
		log.Info("create post error", zap.Error(err))
		return nil, err
	}

	log.Debug("created post", zap.Int("id", post.ID))
	return post, nil
}

func (s *postService) GetPost(ctx context.Context, id int) (*entities.Post, error) {
	log := FromContext(ctx)

	post, err := s.repo.post.GetPost(id)
	if err != nil {
		log.Info("get post error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}

	log.Debug("got post", zap.Int("id", post.ID))
	return post, nil
}

func (s *postService) GetPosts(ctx context.Context, params dto.GetPostParams) ([]entities.Post, int64, error) {
	log := FromContext(ctx)

	posts, total, err := s.repo.post.GetPosts(params)
	if err != nil {
		log.Info("get posts error", zap.Error(err))
		return nil, 0, err
	}

	log.Debug("got posts", zap.Int64("total", total))
	return posts, total, nil
}

func (s *postService) UpdatePost(ctx context.Context, id int, name, contents string) (*entities.Post, error) {
	log := FromContext(ctx)

	post, err := s.repo.post.UpdatePost(id, name, contents)
	if err != nil {
		log.Info("update post error", zap.Error(err))
		return nil, err
	}

	log.Debug("updated post", zap.Int("id", post.ID))
	return post, nil
}

func (s *postService) DeletePost(ctx context.Context, id int) (int, error) {
	log := FromContext(ctx)

	_, err := s.repo.post.DeletePost(id)
	if err != nil {
		log.Info("delete post error", zap.Error(err))
		return 0, err
	}

	log.Debug("deleted post", zap.Int("id", id))
	return id, nil
}
