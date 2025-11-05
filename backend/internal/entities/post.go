package entities

import (
	"time"
)

type Post struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name string `gorm:"not null"`
	Contents string
	AuthorID User `gorm:"foreignKey:User"`
}
