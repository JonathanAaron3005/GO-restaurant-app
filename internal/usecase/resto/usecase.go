package resto

import "github.com/JonathanAaron3005/go-restaurant-app/internal/model"

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	//Order()
	//GetOrderInfo()
}
