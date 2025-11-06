package repository

import (
	"github.com/SemgaTeam/blog/internal/config"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"fmt"
)

func NewPostgresConnection(conf *config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password, 
		conf.Name, 
		conf.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})	
	if err != nil {
		return nil, err
	}

	return db, nil
}
