package handlers

import (
	"context"

	"example.com/mod/order/internal/dto"
	"example.com/mod/order/internal/service"
	order_v1 "example.com/mod/proto/gen/go"
)

type OrderServer struct {
	order_v1.UnimplementedOrderServiceServer
	Repo *service.PostgresOrderRepository
}

func (o *OrderServer) CreateOrderHandler(ctx context.Context, req *order_v1.CreateOrderRequest) (*order_v1.CreateOrderResponse, error) {
	err := o.Repo.CreateOrder(ctx, &dto.Order{
		Name:  req.Name,
		Price: int64(req.Price),
		Count: int32(req.Count),
	})
	return &order_v1.CreateOrderResponse{}, err
}

func (o *OrderServer) UpdateOrderHandler(ctx context.Context, req *order_v1.UpdateOrderRequest) (*order_v1.UpdateOrderResponse, error) {
	err := o.Repo.UpdateOrder(ctx, &dto.Order{
		Name:  req.UpdateOrder.Name,
		Price: req.UpdateOrder.Price,
		Count: req.UpdateOrder.Count,
	})
	return &order_v1.UpdateOrderResponse{}, err
}

func (o *OrderServer) DeleteOrderHandler(ctx context.Context, req *order_v1.DeleteOrderRequest) (*order_v1.DeleteOrderResponse, error) {
	err := o.Repo.DeleteOrder(ctx, req.OrderId)
	return &order_v1.DeleteOrderResponse{}, err
}

func (o *OrderServer) GetOrderHandler(ctx context.Context, req *order_v1.GetOrderRequest) (*order_v1.GetOrderResponse, error) {
	c, err := o.Repo.GetOrder(ctx, req.OrderId)
	if err != nil {
		return &order_v1.GetOrderResponse{}, nil
	}

	return &order_v1.GetOrderResponse{
		OrderName: &order_v1.Order{
			Name:  c.Name,
			Price: c.Price,
			Count: c.Count,
		},
	}, nil
}
