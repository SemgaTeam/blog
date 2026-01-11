package service

import (
	"github.com/SemgaTeam/blog/internal/entities"
	"github.com/SemgaTeam/blog/internal/dto"
)

type PostRepository interface {
	CreatePost(string, string, int) (*entities.Post, error)
	GetPost(int) (*entities.Post, error)
	GetPosts(dto.GetPostParams) ([]entities.Post, int64, error)
	UpdatePost(int, string, string) (*entities.Post, error)
	DeletePost(int) (int, error)
}

type HashRepository interface {
	HashPassword(string) (string, error)
	IsPasswordValid(string, string) bool
}

type TokenRepository interface {
	GenerateAndSignToken(entities.Claims) (*entities.AuthToken, error) 
}

type UserRepository interface {
	CreateUser(string, string) (*entities.User, error)
	GetUserById(int) (*entities.User, error)
	GetUserByName(string) (*entities.User, error)
	GetUsers(dto.GetUserParams) ([]entities.User, int64, error)
	UpdateUser(int, string, string) (*entities.User, error)
	DeleteUser(int) (int, error)
}
