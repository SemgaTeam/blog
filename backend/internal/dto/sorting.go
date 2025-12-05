package dto

import (
	e "github.com/SemgaTeam/blog/internal/error"
	"encoding/json"
	"strings"
)

type Sorting struct {
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
}

func (s *Sorting) UnmarshalParam(str string) error {
	if err := json.Unmarshal([]byte(str), s); err != nil {
		return e.ErrInvalidQueryParam
	}

	s.SortField = strings.TrimSpace(
		strings.ToLower(s.SortField),
	)

	s.SortOrder = strings.TrimSpace(
		strings.ToLower(s.SortOrder),
	)

	if s.SortOrder != "asc" && s.SortOrder != "desc" && s.SortOrder != "" {
		return e.ErrInvalidQueryParam
	}

	if s.SortOrder == "" {
		s.SortOrder = "asc" // default value
	}

	return nil
}
