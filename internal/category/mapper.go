package category

import "github.com/bste101/finance-tracker/db/sqlc"

func ToResponse(c sqlc.Category) CategoryResponse {
	return CategoryResponse{
		ID:     c.ID,
		UserID: c.UserID,
		Name:   c.Name,
		Type:   c.Type,
		Color:  c.Color.String,
		Icon:   c.Icon.String,
	}
}
