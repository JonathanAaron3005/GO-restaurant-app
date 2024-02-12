package order

import (
	"context"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockOrderRepository -destination=../../mocks/order_repository_mock.go -source=repository.go

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
	GetAllOrders(ctx context.Context) ([]model.Order, error)
}
