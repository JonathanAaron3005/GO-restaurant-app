package menu

import (
	"context"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
)

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
	CreateMenu(ctx context.Context, menu model.MenuItem) (model.MenuItem, error)
}
