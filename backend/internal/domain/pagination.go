package domain

import (
	e "github.com/SemgaTeam/blog/internal/error"
)

type pagination struct {
	Page    int
	PerPage int
}

func NewPagination(page, perPage int) (*pagination, error) {
	if page < 0 || perPage <= 0 {
		return nil, e.ErrInvalidPagination
	}

	return &pagination{
		Page:    page,
		PerPage: perPage,
	}, nil
}
