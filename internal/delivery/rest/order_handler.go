package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *handler) Order(c echo.Context) error {
	var request model.OrderMenuRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	orderData, err := h.restoUsecase.Order(request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
