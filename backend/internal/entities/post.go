package entities

import (
	"github.com/SemgaTeam/blog/internal/dto"
	"time"
)

type Post struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name string `gorm:"not null"`
	Contents string
	AuthorID int
	User User `gorm:"foreignKey:AuthorID"`
}

func (p *Post) ToDTO() dto.Post {
	return dto.Post{
		ID: p.ID,
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
		UpdatedAt: p.UpdatedAt.Format(time.RFC3339),
		Name: p.Name,
		Contents: p.Contents,
		AuthorID: p.AuthorID,
	}
}
