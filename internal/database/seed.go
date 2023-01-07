package database

import (
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model/constants"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})
	drinkMenu := []model.MenuItem{
		{
			Name:      "Es teh",
			OrderCode: "es_teh",
			Price:     5000,
			Type:      constants.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     5000,
			Type:      constants.MenuTypeDrink,
		},
		{
			Name:      "Jus Mangga",
			OrderCode: "jus_mangga",
			Price:     5000,
			Type:      constants.MenuTypeDrink,
		},
	}

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constants.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constants.MenuTypeFood,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
	db.AutoMigrate(&model.MenuItem{})
}
