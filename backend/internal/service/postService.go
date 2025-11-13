package service 

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	"github.com/SemgaTeam/blog/internal/repository"
	"github.com/SemgaTeam/blog/internal/log"
	"go.uber.org/zap"
)

type PostService interface {
	CreatePost(string, string, int) (*entities.Post, error)
	GetPost(int) (*entities.Post, error)
	GetPosts(dto.GetPostParams) ([]entities.Post, int64, error)
	UpdatePost(int, string, string) (*entities.Post, error)
	DeletePost(int) (int, error)
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

func (s *postService) CreatePost(name, contents string, authorId int) (*entities.Post, error) {
	post, err := s.repo.post.CreatePost(name, contents, authorId)
	if err != nil {
		log.Log.Info("create post error", zap.Error(err))
		return nil, err
	}

	log.Log.Debug("created post", zap.Int("id", post.ID))
	return post, nil
}

func (s *postService) GetPost(id int) (*entities.Post, error) {
	post, err := s.repo.post.GetPost(id)
	if err != nil {
		log.Log.Info("get post error", zap.Error(err), zap.Int("id", id))
		return nil, err
	}

	log.Log.Debug("got post", zap.Int("id", post.ID))
	return post, nil
}

func (s *postService) GetPosts(params dto.GetPostParams) ([]entities.Post, int64, error) {
	posts, total, err := s.repo.post.GetPosts(params)
	if err != nil {
		log.Log.Info("get posts error", zap.Error(err))
		return nil, 0, err
	}

	log.Log.Debug("got posts", zap.Int64("total", total))
	return posts, total, nil
}

func (s *postService) UpdatePost(id int, name, contents string) (*entities.Post, error) {
	post, err := s.repo.post.UpdatePost(id, name, contents)
	if err != nil {
		log.Log.Info("update post error", zap.Error(err))
		return nil, err
	}

	log.Log.Debug("updated post", zap.Int("id", post.ID))
	return post, nil
}

func (s *postService) DeletePost(id int) (int, error) {
	_, err := s.repo.post.DeletePost(id)
	if err != nil {
		log.Log.Info("delete post error", zap.Error(err))
		return 0, err
	}

	log.Log.Debug("deleted post", zap.Int("id", id))
	return id, nil
}
