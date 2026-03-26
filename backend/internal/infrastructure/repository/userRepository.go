package repository

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"
	"github.com/SemgaTeam/blog/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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

func (r *UserRepository) CreateUser(name, password string) (*entities.User, error) {
	user := entities.User{
		Name:     name,
		Password: password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, e.ErrUserAlreadyExists
		} else {
			return nil, e.Unknown(err)
		}
	}

	return &user, nil
}

func (r *UserRepository) GetUserById(id int) (*entities.User, error) {
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

func (r *UserRepository) GetUserByName(name string) (*entities.User, error) {
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

func (r *UserRepository) GetUsers(params dto.GetUserParams) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	q := r.db.Model(&entities.User{})

	if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	}

	if params.Name != "" {
		q = q.Where("name LIKE ?", "%"+params.Name+"%")
	}

	allowedSortingFields := []string{"id", "name", "created_at"}
	if err := utils.HandleSorting(q, params.Sorting, allowedSortingFields); err != nil {
		return nil, 0, err
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

func (r *UserRepository) UpdateUser(id int, name, password string) (*entities.User, error) {
	user := entities.User{
		ID:       id,
		Name:     name,
		Password: password,
	}

	if err := r.db.
		Clauses(clause.Returning{}).
		Updates(&user).
		Scan(&user).Error; err != nil {
		return nil, e.Unknown(err)
	}

	return &user, nil
}

func (r *UserRepository) DeleteUser(id int) (int, error) {
	user := entities.User{
		ID: id,
	}

	res := r.db.Delete(&user)

	if err := res.Error; err != nil {
		return 0, e.Unknown(err)
	}

	if res.RowsAffected == 0 {
		return 0, e.ErrUserNotFound
	}

	return id, nil
}
