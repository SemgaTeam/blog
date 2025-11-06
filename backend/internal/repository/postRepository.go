package repository

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(string, string, int) (*entities.Post, error)
	GetPost(int) (*entities.Post, error)
	UpdatePost(int, string, string) (*entities.Post, error)
	DeletePost(int) (int, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) CreatePost(name, contents string, authorId int) (*entities.Post, error) {
	post := entities.Post{
		Name: name,
		Contents: contents,
		AuthorID: authorId,
	}

	if err := r.db.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) GetPost(id int) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.Where("id = ?", id).Find(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) UpdatePost(id int, name, contents string) (*entities.Post, error) {
	post := entities.Post{
		ID: id,
		Name: name,
		Contents: contents,
	}

	if err := r.db.Updates(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) DeletePost(id int) (int, error) {
	post := entities.Post{
		ID: id,
	}

	if err := r.db.Delete(post).Error; err != nil {
		return 0, err
	}

	return id, nil
}
