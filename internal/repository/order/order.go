package order

import (
	"context"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/tracing"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}

func (or *orderRepo) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()

	if err := or.db.WithContext(ctx).Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (or *orderRepo) GetOrderInfo(ctx context.Context, orderID string) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()

	var orderInfo model.Order

	if err := or.db.WithContext(ctx).Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&orderInfo).Error; err != nil {
		return orderInfo, err
	}

	return orderInfo, nil
}

func (or *orderRepo) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	var orderData []model.Order

	if err := or.db.WithContext(ctx).Preload("ProductOrders").Find(&orderData).Error; err != nil {
		return orderData, err
	}

	return orderData, nil
}
