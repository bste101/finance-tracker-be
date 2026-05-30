package user

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

func (r *Repository) GetByID(ctx context.Context, id int64) (sqlc.GetUserByIDRow, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *Repository) Update(ctx context.Context, id int64, name string) (sqlc.UpdateUserRow, error) {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
	})
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id)
}
