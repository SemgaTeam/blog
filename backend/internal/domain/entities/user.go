package entities

import (
	"github.com/SemgaTeam/blog/internal/dto"
	e "github.com/SemgaTeam/blog/internal/error"

	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Name      string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	IsAdmin   bool      `gorm:"default:false"`
}

func NewUser(name, password string) (*User, error) {
	if name == "" || password == "" {
		return nil, e.ErrInvalidUser
	}

	return &User{
		Name:     name,
		Password: password,
	}, nil
}

func UpdateUser(id int, name, password string) (*User, error) {
	if name == "" || password == "" {
		return nil, e.ErrInvalidUser
	}

	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}, nil
}

func (u *User) ToDTO() dto.User {
	return dto.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		Name:      u.Name,
		IsAdmin:   u.IsAdmin,
	}
}
