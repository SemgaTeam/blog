package repository

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/utils"
	"gorm.io/gorm"

	"errors"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Save(post *entities.Post) error {
	if err := r.db.Save(post).Error; err != nil {
		return e.Unknown(err)
	}

	return nil
}

func (r *PostRepository) ById(id int) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.Where("id = ?", id).Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrPostNotFound
		} else {
			return nil, e.Unknown(err)
		}
	}

	return &post, nil
}

func (r *PostRepository) ByParams(params dto.GetPostParams) ([]entities.Post, int64, error) {
	var posts []entities.Post
	var total int64

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

	if params.Sorting.SortField != "" && params.Sorting.SortOrder != "" {
		q = q.Order(params.Sorting.SortField + " " + params.Sorting.SortOrder)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, e.Unknown(err)
	}

	utils.HandlePagination(q, params.Pagination)

	res := q.Find(&posts)

	if err := res.Error; err != nil {
		return nil, 0, e.Unknown(err)
	}

	if total == 0 {
		return nil, 0, e.ErrPostNotFound
	}

	return posts, total, nil
}

func (r *PostRepository) Delete(id int) error {
	post := entities.Post{
		ID: id,
	}

	res := r.db.Delete(post)

	if err := res.Error; err != nil {
		return e.Unknown(err)
	}

	if res.RowsAffected == 0 {
		return e.ErrPostNotFound
	}

	return nil
}
