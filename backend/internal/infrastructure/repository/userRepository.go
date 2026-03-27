package repository

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/utils"
	"gorm.io/gorm"

	"errors"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(user *entities.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return e.Unknown(err)
	}

	return nil
}

func (r *UserRepository) ById(id int) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserNotFound
		} else {
			return nil, e.Unknown(err)
		}
	}

	return &user, nil
}

func (r *UserRepository) ByName(name string) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserNotFound
		} else {
			return nil, e.Unknown(err)
		}
	}

	return &user, nil
}

func (r *UserRepository) ByParams(params dto.GetUserParams) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	q := r.db.Model(&entities.User{})

	if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	}

	if params.Name != "" {
		q = q.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Sorting.SortField != "" && params.Sorting.SortOrder != "" {
		q = q.Order(params.Sorting.SortField + " " + params.Sorting.SortOrder)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, e.Unknown(err)
	}

	utils.HandlePagination(q, params.Pagination)

	res := q.Find(&users)

	if err := res.Error; err != nil {
		return nil, 0, e.Unknown(err)
	}

	if total == 0 {
		return nil, 0, e.ErrUserNotFound
	}

	return users, total, nil
}

func (r *UserRepository) Delete(id int) error {
	user := entities.User{
		ID: id,
	}

	res := r.db.Delete(&user)

	if err := res.Error; err != nil {
		return e.Unknown(err)
	}

	if res.RowsAffected == 0 {
		return e.ErrUserNotFound
	}

	return nil
}
