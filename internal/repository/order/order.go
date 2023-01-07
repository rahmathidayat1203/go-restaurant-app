package order

import (
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &OrderRepo{
		db: db,
	}
}

func (or *OrderRepo) CreateOrder(order model.Order) (model.Order, error) {
	if err := or.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
func (or *OrderRepo) GetOrderInfo(orderID string) (model.Order, error) {
	var data model.Order

	if err := or.db.Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, nil
	}
	return data, nil
}
