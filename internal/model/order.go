package model

type OrderStatus string

type Order struct {
	ID            string `gorm:"primaryKey"`
	Status        OrderStatus
	ProductOrders []ProductOrder
}

type ProductOrderStatus string

type ProductOrder struct {
	ID         string `gorm:"primaryKey"`
	OrderID    string //foreign key
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     ProductOrderStatus
}

//model utk database dipisah sama request, yg nanti bakal di convert

type OrderMenuProductRequest struct {
	OrderCode string
	Quantity  int
}

type OrderMenuRequest struct {
	MenuProductRequests []OrderMenuProductRequest
}

type GetOrderInfoRequest struct {
	OrderId string
}
