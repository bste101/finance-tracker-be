package auth

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

func (r *Repository) CreateUser(
	ctx context.Context,
	arg sqlc.CreateUserParams,
) (sqlc.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *Repository) GetUserByEmail(
	ctx context.Context,
	email string,
) (sqlc.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}
