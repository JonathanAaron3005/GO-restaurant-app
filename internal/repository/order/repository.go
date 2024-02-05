package order

import (
	"context"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
)

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
	GetAllOrders(ctx context.Context) ([]model.Order, error)
}
