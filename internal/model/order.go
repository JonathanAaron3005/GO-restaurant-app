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
