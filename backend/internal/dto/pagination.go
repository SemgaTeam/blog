package dto

import (
	e "github.com/SemgaTeam/blog/internal/error"
	"encoding/json"
)

type Pagination struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
}

func (p *Pagination) UnmarshalParam(str string) error {
	if err := json.Unmarshal([]byte(str), p); err != nil {
		return e.ErrInvalidQueryParam
	}

	return nil
}
