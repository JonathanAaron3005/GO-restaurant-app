package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenuList(c echo.Context) error {
	menuType := c.FormValue("menu_type") //food or drink

	menuData, err := h.restoUsecase.GetMenuList(menuType)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func (h *handler) AddNewMenu(c echo.Context) error {
	var menuData model.MenuItem

	err := json.NewDecoder(c.Request().Body).Decode(&menuData)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	createdMenu, err := h.restoUsecase.AddNewMenu(menuData)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": createdMenu,
	})
}
