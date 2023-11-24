package order

import (
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
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

func (or *orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	if err := or.db.Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (or *orderRepo) GetOrderInfo(orderID string) (model.Order, error) {
	var orderInfo model.Order

	if err := or.db.Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&orderInfo).Error; err != nil {
		return orderInfo, err
	}

	return orderInfo, nil
}

func (or *orderRepo) GetAllOrders() ([]model.Order, error) {
	var orderData []model.Order

	if err := or.db.Preload("ProductOrders").Find(&orderData).Error; err != nil {
		return orderData, err
	}

	return orderData, nil
}
