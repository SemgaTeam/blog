package application

import (
	"github.com/SemgaTeam/blog/internal/domain/entities"
	"github.com/SemgaTeam/blog/internal/dto"

	"context"
)

type Service interface {
	AuthService
	PostService
	UserService
}

type AuthService interface {
	LogIn(ctx context.Context, name, password string) (*entities.AuthToken, *entities.AuthToken, error)
	SignIn(ctx context.Context, name string, password string) (*entities.AuthToken, *entities.AuthToken, error)
	RefreshTokens(ctx context.Context, userId int, isAdmin bool) (*entities.AuthToken, *entities.AuthToken, error)
	GetMe(ctx context.Context, userId int) (*entities.User, error)
}

type PostService interface {
	CreatePost(ctx context.Context, name, contents string, authorId int) (*entities.Post, error)
	GetPost(ctx context.Context, id int) (*entities.Post, error)
	GetPosts(ctx context.Context, params dto.GetPostParams) ([]entities.Post, int64, error)
	UpdatePost(ctx context.Context, id int, name string, contents string) (*entities.Post, error)
	DeletePost(ctx context.Context, id int) (int, error)
}

type UserService interface {
	CreateUser(ctx context.Context, name string, password string) (*entities.User, error)
	GetUserById(ctx context.Context, id int) (*entities.User, error)
	GetUsers(ctx context.Context, params dto.GetUserParams) ([]entities.User, int64, error)
	UpdateUser(ctx context.Context, id int, name string, password string) (*entities.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}
