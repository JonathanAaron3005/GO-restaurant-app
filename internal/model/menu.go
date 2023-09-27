package model

type MenuType string //alias

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int64
	Type      MenuType
}
