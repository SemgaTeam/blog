package db

import (
	"github.com/SemgaTeam/blog/internal/config"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/pressly/goose/v3"
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
)

func NewPostgresConnection(conf *config.Postgres) (*gorm.DB, error) {
	dsn := postgresDSN(conf)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})	
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(conf *config.Postgres, migrationsPath string) error {
	dsn := postgresDSN(conf)
	sqlDb, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	if err := goose.Up(sqlDb, migrationsPath); err != nil {
		return err
	}

	return nil
}

func postgresDSN(conf *config.Postgres) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password, 
		conf.Name, 
		conf.Port,
	)

}
