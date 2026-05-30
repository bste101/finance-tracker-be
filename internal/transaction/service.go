package transaction

import (
	"context"

	"github.com/bste101/finance-tracker/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, userID int64, req CreateTransactionRequest) (TransactionResponse, error) {
	t, err := s.repo.Create(ctx, sqlc.CreateTransactionParams{
		UserID:          userID,
		CategoryID:      req.CategoryID,
		Type:            req.Type,
		Amount:          req.Amount,
		Note:            pgtype.Text{String: req.Note, Valid: req.Note != ""},
		TransactionDate: pgtype.Date{Time: req.TransactionDate, Valid: true},
	})
	if err != nil {
		return TransactionResponse{}, err
	}

	return ToResponse(t), nil
}

func (s *Service) FindAll(ctx context.Context, userID int64) ([]ListTransactionResponse, error) {
	rows, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]ListTransactionResponse, 0, len(rows))
	for _, row := range rows {
		res = append(res, ToListResponse(row))
	}

	return res, nil
}

func (s *Service) Update(ctx context.Context, userID int64, id int64, req UpdateTransactionRequest) (TransactionResponse, error) {
	t, err := s.repo.Update(ctx, sqlc.UpdateTransactionParams{
		ID:              id,
		UserID:          userID,
		CategoryID:      req.CategoryID,
		Type:            req.Type,
		Amount:          req.Amount,
		Note:            pgtype.Text{String: req.Note, Valid: req.Note != ""},
		TransactionDate: pgtype.Date{Time: req.TransactionDate, Valid: true},
	})
	if err != nil {
		return TransactionResponse{}, err
	}

	return ToResponse(t), nil
}

func (s *Service) Delete(ctx context.Context, userID int64, id int64) error {
	return s.repo.Delete(ctx, sqlc.DeleteTransactionParams{
		ID:     id,
		UserID: userID,
	})
}
