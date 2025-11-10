package repository

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"errors"
	"strings"
)

type PostRepository interface {
	CreatePost(string, string, int) (*entities.Post, error)
	GetPost(int) (*entities.Post, error)
	GetPosts(dto.GetPostParams) ([]entities.Post, int64, error)
	UpdatePost(int, string, string) (*entities.Post, error)
	DeletePost(int) (int, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
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
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return nil, e.BadRequest(err, "invalid request body")
		}	else {
			return nil, e.Internal(err)
		}
	}

	return &post, nil
}

func (r *postRepository) GetPost(id int) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.Where("id = ?", id).Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrPostNotFound
		} else {
			return nil, e.Internal(err)
		}
	}

	return &post, nil
}

func (r *postRepository) GetPosts(params dto.GetPostParams) ([]entities.Post, int64, error) {
	var posts []entities.Post
	var total int64
	params.SortField = strings.TrimSpace(strings.ToLower(params.SortField))
	params.SortOrder = strings.TrimSpace(strings.ToLower(params.SortOrder))

	q := r.db.Model(&entities.Post{})	

	if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	}

	if params.Name != "" {
		q = q.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.AuthorID != 0 {
		q = q.Where("author_id = ?", params.AuthorID)
	}

	allowedSortingFields := []string{"created_at", "updated_at"}
	if err := handleSorting(q, params.SortField, params.SortOrder, allowedSortingFields); err != nil {
		return nil, 0, err
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, e.Internal(err)
	}

	if params.Page != 0 && params.PerPage != 0 {
		q = q.Limit(params.PerPage).Offset((params.Page - 1)*params.PerPage)
	}

	res := q.Find(&posts)

	if err := res.Error; err != nil {
		return nil, 0, e.Internal(err)
	}

	if total == 0 {
		return nil, 0, e.ErrPostNotFound
	}

	return posts, total, nil
}

func (r *postRepository) UpdatePost(id int, name, contents string) (*entities.Post, error) {
	post := entities.Post{
		ID: id,
		Name: name,
		Contents: contents,
	}

	if err := r.db.
							Clauses(clause.Returning{}).
							Updates(&post).
							Scan(&post).Error; 
							err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return nil, e.BadRequest(err, "invalid request body")
		}	else {
			return nil, e.Internal(err)
		}
	}

	return &post, nil
}

func (r *postRepository) DeletePost(id int) (int, error) {
	post := entities.Post{
		ID: id,
	}

	res := r.db.Delete(post)

	if err := res.Error; err != nil {
		return 0, e.Internal(err)
	}

	if res.RowsAffected == 0 {
		return 0, e.ErrPostNotFound
	}

	return id, nil
}
