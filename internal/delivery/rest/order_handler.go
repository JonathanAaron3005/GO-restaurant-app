package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model/constant"
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

	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

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

func (h *handler) GetOrderInfo(c echo.Context) error {
	orderID := c.Param("orderID")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUsecase.GetOrderInfo(model.GetOrderInfoRequest{
		UserID:  userID,
		OrderId: orderID,
	})
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

func (h *handler) getAllOrdersInfo(c echo.Context) error {
	orderData, err := h.restoUsecase.GetAllOrdersInfo()

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
