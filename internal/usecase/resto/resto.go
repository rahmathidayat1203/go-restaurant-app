package resto

import (
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/repository/menu"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUseCase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}
func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}
