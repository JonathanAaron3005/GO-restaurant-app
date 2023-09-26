package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie ayam",
			OrderCode: "bakmie_ayam",
			price:     31200,
		},
		{
			Name:      "Nasi goreng",
			OrderCode: "nasi_goreng",
			price:     25000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Badak",
			OrderCode: "badak",
			price:     8000,
		},
		{
			Name:      "Ice tea",
			OrderCode: "ice_tea",
			price:     3000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()

	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
