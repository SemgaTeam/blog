package service 

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/repository"
)

type PostService interface {
	CreatePost(string, string, int) (*entities.Post, error)
	GetPost(int) (*entities.Post, error)
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
	return s.repo.post.CreatePost(name, contents, authorId)
}

func (s *postService) GetPost(id int) (*entities.Post, error) {
	return s.repo.post.GetPost(id)
}

func (s *postService) UpdatePost(id int, name, contents string) (*entities.Post, error) {
	return s.repo.post.UpdatePost(id, name, contents)
}

func (s *postService) DeletePost(id int) (int, error) {
	return s.repo.post.DeletePost(id)
}
