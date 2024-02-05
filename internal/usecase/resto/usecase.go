package resto

import (
	"context"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
)

type Usecase interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	AddNewMenu(ctx context.Context, menu model.MenuItem) (model.MenuItem, error)
	Order(ctx context.Context, req model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(ctx context.Context, req model.GetOrderInfoRequest) (model.Order, error)
	GetAllOrdersInfo(ctx context.Context) ([]model.Order, error)
	RegisterUser(ctx context.Context, req model.RegisterRequest) (model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}
