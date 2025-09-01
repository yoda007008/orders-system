package service

import (
	"context"
	"fmt"

	"example.com/mod/order/internal/dto"
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
func (p *PostgresOrderRepository) CreateOrder(ctx context.Context, c *dto.Order) error {
	_, err := p.db.Exec(ctx, "INSERT INTO order(name, price, count, order_id) VALUES ($1, $2, $3)")
	return err
}
