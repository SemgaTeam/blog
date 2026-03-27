package domain

import (
	e "github.com/SemgaTeam/blog/internal/error"
)

var validUsersortingFields = map[string]struct{}{
	"id":         {},
	"name":       {},
	"created_at": {},
}

var validPostsortingFields = map[string]struct{}{
	"id":         {},
	"name":       {},
	"created_at": {},
	"updated_at": {},
}

type sorting struct {
	SortField string
	SortOrder string
}

func NewUserSorting(field, order string) (*sorting, error) {
	if _, ok := validUsersortingFields[field]; !ok {
		return nil, e.ErrInvalidSortingField
	}

	if (order != "asc") && (order != "desc") {
		return nil, e.ErrInvalidSortOrder
	}

	return &sorting{
		SortField: field,
		SortOrder: order,
	}, nil
}

func NewPostSorting(field, order string) (*sorting, error) {
	if _, ok := validPostsortingFields[field]; !ok {
		return nil, e.ErrInvalidSortingField
	}

	if (order != "asc") && (order != "desc") {
		return nil, e.ErrInvalidSortOrder
	}

	return &sorting{
		SortField: field,
		SortOrder: order,
	}, nil
}
