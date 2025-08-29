package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderRepository struct {
	db *pgxpool.Pool
}

func NewPostgresOrderRepository(connsStr string) (*PostgresOrderRepository, error) {
	db, err := pgxpool.New(context.Background(), connsStr)
	if err != nil {
		return nil, fmt.Errorf("failed connection database")
	}

	return &PostgresOrderRepository{
		db: db,
	}, nil
}

// todo bisnes logic (create, read, update, delete)
