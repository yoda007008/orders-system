package service

import (
	"context"
	"fmt"
	"log/slog"

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
	_, err := p.db.Exec(ctx, "INSERT INTO orders (order_name, order_price, order_count) VALUES ($1, $2, $3)", c.Name, c.Price, c.Count)
	slog.Info("INSERT")
	return err
}

func (p *PostgresOrderRepository) UpdateOrder(ctx context.Context, c *dto.Order) error {
	_, err := p.db.Exec(ctx, `UPDATE orders SET order_name = $1, order_price = $2, order_count = $3`, c.Name, c.Count, c.Price)
	slog.Info("UPDATE")
	return err
}

func (p *PostgresOrderRepository) DeleteOrder(ctx context.Context, order_id int32) error {
	_, err := p.db.Exec(ctx, `DELETE FROM orders WHERE order_id = $1`, order_id)
	slog.Info("DELETE")
	return err
}

func (p *PostgresOrderRepository) GetOrder(ctx context.Context, order_id int32) (*dto.Order, error) {
	row := p.db.QueryRow(ctx, `SELECT order_name, order_price, order_count FROM order WHERE order_id = $1`, order_id)
	var c *dto.Order
	err := row.Scan(&c.Name, &c.Price, &c.Count)
	slog.Info("GET")
	return c, err
}
