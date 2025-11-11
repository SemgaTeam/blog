package entities

import (
	"github.com/SemgaTeam/blog/internal/dto"
	"time"
)

type User struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Name string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func (u *User) ToDTO() dto.User {
	return dto.User{
		ID: u.ID,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		Name: u.Name,
	}
}
