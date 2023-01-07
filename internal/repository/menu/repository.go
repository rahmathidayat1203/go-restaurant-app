package menu

import "github.com/rahmathidayat1203/go-restaurant-app/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
