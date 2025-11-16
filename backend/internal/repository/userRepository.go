package repository

import (
	"github.com/SemgaTeam/blog/internal/entities"
	e "github.com/SemgaTeam/blog/internal/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"errors"
)

type UserRepository interface {
	CreateUser(string, string) (*entities.User, error)
	GetUserById(int) (*entities.User, error)
	GetUserByName(string) (*entities.User, error)
	UpdateUser(int, string, string) (*entities.User, error)
	DeleteUser(int) (int, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(name, password string) (*entities.User, error) {
	user := entities.User{
		Name: name,
		Password: password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, e.ErrUserAlreadyExists
		} else {
			return nil, e.Internal(err)
		}
	}

	return &user, nil
}

func (r *userRepository) GetUserById(id int) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserNotFound
		} else {
			return nil, e.Internal(err)
		}
	}

	return &user, nil
}

func (r *userRepository) GetUserByName(name string) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserNotFound
		} else {
			return nil, e.Internal(err)
		}
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(id int, name, password string) (*entities.User, error) {
	user := entities.User{
		ID: id,
		Name: name,
		Password: password,
	}

	if err := r.db.
								Clauses(clause.Returning{}).
								Updates(&user).
								Scan(&user).Error;
						err != nil {
		return nil, e.Internal(err)
	}

	return &user, nil
}

func (r *userRepository) DeleteUser(id int) (int, error) {
	user := entities.User{
		ID: id,
	}

	res := r.db.Delete(&user)

	if err := res.Error; err != nil {
		return 0, e.Internal(err)
	}

	if res.RowsAffected == 0 {
		return 0, e.NotFound(e.ErrUserNotFound, "user not found")
	}

	return id, nil
}
