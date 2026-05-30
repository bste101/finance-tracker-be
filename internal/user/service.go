package user

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Me(ctx context.Context, userID int64) (*ProfileResponse, error) {
	u, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &ProfileResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *Service) Update(ctx context.Context, userID int64, req UpdateProfileRequest) (*ProfileResponse, error) {
	u, err := s.repo.Update(ctx, userID, req.Name)
	if err != nil {
		return nil, err
	}

	return &ProfileResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *Service) Delete(ctx context.Context, userID int64) error {
	return s.repo.Delete(ctx, userID)
}
