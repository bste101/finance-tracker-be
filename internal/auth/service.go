package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/bste101/finance-tracker/db/sqlc"
)

type Service struct {
	q *sqlc.Queries
}

func NewService(q *sqlc.Queries) *Service {
	return &Service{q: q}
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) (AuthResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, err
	}

	user, err := s.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return AuthResponse{}, err
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{AccessToken: token}, nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (AuthResponse, error) {
	user, err := s.q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return AuthResponse{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return AuthResponse{}, errors.New("invalid email or password")
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{AccessToken: token}, nil
}

func generateToken(userID int64) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr, secret, issuer string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("failed to parse claims")
}
