package resto

import (
	"errors"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model/constant"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/repository/menu"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/repository/order"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/repository/user"
	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
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

func (r *restoUsecase) GetAllOrdersInfo() ([]model.Order, error) {
	return r.orderRepo.GetAllOrders()
}

func (r *restoUsecase) RegisterUser(req model.RegisterRequest) (model.User, error) {
	userRegistered, err := r.userRepo.CheckRegistered(req.Username)

	if err != nil {
		return model.User{}, err
	}

	if userRegistered {
		return model.User{}, errors.New("User already registered")
	}

	userHash, err := r.userRepo.GenerateUserHash(req.Password)

	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(model.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Hash:     userHash,
	})

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (r *restoUsecase) Login(req model.LoginRequest) (model.UserSession, error) {
	userData, err := r.userRepo.GetUserData(req.Username)
	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := r.userRepo.VerifyLogin(req.Username, req.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}

	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := r.userRepo.CreateUserSession(userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}
