package category

import (
	"context"

	"github.com/bste101/finance-tracker/db/sqlc"
)

type Repository struct {
	q *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) Create(ctx context.Context, arg sqlc.CreateCategoryParams) (sqlc.Category, error) {
	return r.q.CreateCategory(ctx, arg)
}

func (r *Repository) FindAll(ctx context.Context, userID int64) ([]sqlc.Category, error) {
	return r.q.GetCategories(ctx, userID)
}

func (r *Repository) Update(ctx context.Context, arg sqlc.UpdateCategoryParams) (sqlc.Category, error) {
	return r.q.UpdateCategory(ctx, arg)
}

func (r *Repository) Delete(ctx context.Context, arg sqlc.DeleteCategoryParams) error {
	return r.q.DeleteCategory(ctx, arg)
}
