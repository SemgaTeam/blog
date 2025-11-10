package repository 

import (
	e "github.com/SemgaTeam/blog/internal/error"
	"gorm.io/gorm"

	"strings"
)

func handleSorting(q *gorm.DB, sortField, sortOrder string, allowedFields []string) error {
	sortField = strings.TrimSpace(strings.ToLower(sortField))
	sortOrder = strings.TrimSpace(strings.ToLower(sortOrder))

	if sortField == "" { 
		return nil
	}

	if sortOrder == "" {
		sortOrder = "asc" // default value
	}

	allowed := false
	for _, allowedField := range allowedFields {
		if sortField == allowedField {
			allowed = true
		}
	}

	if !allowed {
		return e.BadRequest(nil, "sort field is not allowed")
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		return e.BadRequest(nil, "sort order is invalid")
	}

	q = q.Order(sortField + " " + sortOrder)

	return nil
}
