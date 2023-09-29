package resto

import (
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model/constant"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/repository/menu"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/repository/order"
	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
	}
}

func (r *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUsecase) AddNewMenu(menu model.MenuItem) (model.MenuItem, error) {
	return r.menuRepo.CreateMenu(menu)
}

func (r *restoUsecase) Order(req model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(req.OrderProducts))

	for i, orderProduct := range req.OrderProducts {
		//mastiin bahwa product yang di order ada di db
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menuData.Price * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProccessed,
		ProductOrders: productOrderData,
		ReferenceID:   req.ReferenceID,
	}

	createdOrder, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createdOrder, nil

}

func (r *restoUsecase) GetOrderInfo(req model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(req.OrderId)

	if err != nil {
		return orderData, err
	}

	return orderData, nil
}
