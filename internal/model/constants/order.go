package constants

import "github.com/rahmathidayat1203/go-restaurant-app/internal/model"

const (
	OrderStatusProcessed model.OrderStatus = "processed"
	OrderStatusFinished  model.OrderStatus = "finished"
	OrderStatusFailed    model.OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing model.OrderStatus = "preparing"
	ProductOrderStatusFinished  model.OrderStatus = "finished"
)
