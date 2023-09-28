package menu

import "github.com/JonathanAaron3005/go-restaurant-app/internal/model"

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
