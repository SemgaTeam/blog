package domain

import (
	e "github.com/SemgaTeam/blog/internal/error"
)

type SortField string
type SortOrder string

var (
	IDField        SortField = "id"
	NameField      SortField = "name"
	CreatedAtField SortField = "created_at"
	UpdatedAtField SortField = "updated_at"
)

var validUsersortingFields = map[SortField]struct{}{
	IDField:        {},
	NameField:      {},
	CreatedAtField: {},
}

var validPostsortingFields = map[SortField]struct{}{
	IDField:        {},
	NameField:      {},
	CreatedAtField: {},
	UpdatedAtField: {},
}

type sorting struct {
	SortField SortField `json:"sortField"`
	SortOrder SortOrder `json:"sortOrder"`
}

func NewUsersorting(field SortField, order SortOrder) (*sorting, error) {
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

func NewPostSorting(field SortField, order SortOrder) (*sorting, error) {
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
