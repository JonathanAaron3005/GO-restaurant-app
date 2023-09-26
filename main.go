package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	seedDB()
	e := echo.New()

	e.GET("/menu", getMenu)

	e.Logger.Fatal(e.Start(":14045"))
}

const (
	dbAddress = "host=localhost port=5432 user=postgres password=ja3005 dbname=go_resto_app sslmode=disable"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int64
	Type      MenuType
}

type MenuType string //alias

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie ayam",
			OrderCode: "bakmie_ayam",
			Price:     31200,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Nasi goreng",
			OrderCode: "nasi_goreng",
			Price:     25000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Badak",
			OrderCode: "badak",
			Price:     8000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Ice tea",
			OrderCode: "ice_tea",
			Price:     3000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu)
	fmt.Println(drinkMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err) //utk stop smua coding
	}

	//db.AutoMigrate(&MenuItem{})

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type") //food or drink

	var menuData []MenuItem

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
