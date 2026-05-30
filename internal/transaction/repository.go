package transaction

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

func (r *Repository) Create(ctx context.Context, arg sqlc.CreateTransactionParams) (sqlc.Transaction, error) {
	return r.q.CreateTransaction(ctx, arg)
}

func (r *Repository) FindAll(ctx context.Context, userID int64) ([]sqlc.ListTransactionsRow, error) {
	return r.q.ListTransactions(ctx, userID)
}

func (r *Repository) Update(ctx context.Context, arg sqlc.UpdateTransactionParams) (sqlc.Transaction, error) {
	return r.q.UpdateTransaction(ctx, arg)
}

func (r *Repository) Delete(ctx context.Context, arg sqlc.DeleteTransactionParams) error {
	return r.q.DeleteTransaction(ctx, arg)
}
