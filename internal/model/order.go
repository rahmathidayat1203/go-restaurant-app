package model

type OrderStatus string

type Order struct {
	Status        OrderStatus
	ProductOrders []ProductOrder
}

type ProductOrderStatus string

type ProductOrder struct {
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     ProductOrderStatus
}
