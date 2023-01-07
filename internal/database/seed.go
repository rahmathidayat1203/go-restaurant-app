package database

import (
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	drinkMenu := []model.MenuItem{
		{
			Name:      "Es teh",
			OrderCode: "es_teh",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Jus Mangga",
			OrderCode: "jus_mangga",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
	}

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
	db.AutoMigrate(&model.MenuItem{})
}
