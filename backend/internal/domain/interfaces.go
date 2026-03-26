package domain

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"
)

type UserRepository interface {
	Save(user *entities.User) error
	Delete(id int) error

	ById(id int) (*entities.User, error)
	ByName(name string) (*entities.User, error)
	ByParams(params dto.GetUserParams) ([]entities.User, int64, error)
}

type PostRepository interface {
	Save(post *entities.Post) error
	ById(id int) (*entities.Post, error)
	ByParams(params dto.GetPostParams) ([]entities.Post, int64, error)
	Delete(id int) error
}

type TokenRepository interface {
	GenerateAndSignToken(claims entities.Claims) (*entities.AuthToken, error)
}

type HashRepository interface {
	HashPassword(rawPassword string) (string, error)
	IsPasswordValid(raw string, hash string) bool
}
