package database

import (
	"fmt"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie ayam",
			OrderCode: "bakmie_ayam",
			Price:     31200,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Nasi goreng",
			OrderCode: "nasi_goreng",
			Price:     25000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Badak",
			OrderCode: "badak",
			Price:     8000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Ice tea",
			OrderCode: "ice_tea",
			Price:     3000,
			Type:      constant.MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu)
	fmt.Println(drinkMenu)

	//db.AutoMigrate(&MenuItem{})

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}
