package handlers

import (
	"context"

	"example.com/mod/order/internal/dto"
	"example.com/mod/order/internal/service"
	order_v1 "example.com/mod/proto/gen/go"
)

type OrderServer struct {
	order_v1.UnimplementedOrderServiceServer
	repo *service.PostgresOrderRepository
}

func (o *OrderServer) CreateOrderHandler(ctx context.Context, req *order_v1.CreateOrderRequest) (*order_v1.CreateOrderResponse, error) {
	err := o.repo.CreateOrder(ctx, &dto.Order{
		Name:  req.Name,
		Price: int64(req.Price),
		Count: int32(req.Count),
	})
	return &order_v1.CreateOrderResponse{}, err
}
