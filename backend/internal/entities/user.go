package entities

import (
	"time"
)

type User struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Name string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
