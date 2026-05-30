package transaction

import "github.com/bste101/finance-tracker/db/sqlc"

func ToResponse(t sqlc.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:              t.ID,
		UserID:          t.UserID,
		CategoryID:      t.CategoryID,
		Type:            t.Type,
		Amount:          t.Amount,
		Note:            t.Note.String,
		TransactionDate: t.TransactionDate.Time,
	}
}

func ToListResponse(t sqlc.ListTransactionsRow) ListTransactionResponse {
	return ListTransactionResponse{
		ID:              t.ID,
		CategoryID:      t.CategoryID,
		CategoryName:    t.CategoryName,
		Type:            t.Type,
		Amount:          t.Amount,
		Note:            t.Note.String,
		TransactionDate: t.TransactionDate.Time,
	}
}
