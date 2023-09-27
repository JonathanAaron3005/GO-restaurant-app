package constant

import "github.com/JonathanAaron3005/go-restaurant-app/internal/model"

const (
	OrderStatusProccessed model.OrderStatus = "proccessed"
	OrderStatusFinished   model.OrderStatus = "finished"
	OrderStatusFailed     model.OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing model.ProductOrderStatus = "preparing"
	ProductOrderStatusFinished  model.ProductOrderStatus = "finished"
)
