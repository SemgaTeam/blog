package repository

import (
	"github.com/SemgaTeam/blog/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type HashRepository struct {
	conf *config.Hash
}

func NewHashRepository(conf *config.Hash) *HashRepository {
	return &HashRepository{
		conf: conf,
	}
}

func (r *HashRepository) HashPassword(raw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(raw), r.conf.Cost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (r *HashRepository) IsPasswordValid(rawPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawPassword))
	return err == nil
}
