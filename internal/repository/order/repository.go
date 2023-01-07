package order

import "github.com/rahmathidayat1203/go-restaurant-app/internal/model"

type Repository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderID string) (model.Order, error)
}
