package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
}
