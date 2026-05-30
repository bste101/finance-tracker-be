package category

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

func (s *Service) Create(ctx context.Context, userID int64, req CreateCategoryRequest) (CategoryResponse, error) {
	c, err := s.repo.Create(ctx, sqlc.CreateCategoryParams{
		UserID: userID,
		Name:   req.Name,
		Type:   req.Type,
		Color: pgtype.Text{
			String: req.Color,
			Valid:  req.Color != "",
		},
		Icon: pgtype.Text{
			String: req.Icon,
			Valid:  req.Icon != "",
		},
	})
	if err != nil {
		return CategoryResponse{}, err
	}

	return ToResponse(c), nil
}

func (s *Service) FindAll(ctx context.Context, userID int64) ([]CategoryResponse, error) {
	categories, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]CategoryResponse, 0, len(categories))
	for _, c := range categories {
		res = append(res, ToResponse(c))
	}

	return res, nil
}

func (s *Service) Update(ctx context.Context, userID int64, id int64, req UpdateCategoryRequest) (CategoryResponse, error) {
	c, err := s.repo.Update(ctx, sqlc.UpdateCategoryParams{
		ID:   id,
		Name: req.Name,
		Type: req.Type,
		Color: pgtype.Text{
			String: req.Color,
			Valid:  req.Color != "",
		},
		Icon: pgtype.Text{
			String: req.Icon,
			Valid:  req.Icon != "",
		},
		UserID: userID,
	})
	if err != nil {
		return CategoryResponse{}, err
	}

	return ToResponse(c), nil
}

func (s *Service) Delete(ctx context.Context, userID int64, id int64) error {
	return s.repo.Delete(ctx, sqlc.DeleteCategoryParams{
		ID:     id,
		UserID: userID,
	})
}
