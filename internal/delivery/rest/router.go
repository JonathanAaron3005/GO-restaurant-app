package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	e.GET("/menu", handler.GetMenuList)
	e.POST("/menu", handler.AddNewMenu)

	e.GET("/order", handler.getAllOrdersInfo)
	e.GET("/order/:orderID", handler.GetOrderInfo)
	e.POST("/order", handler.Order)

	e.POST("/user/register", handler.RegisterUser)
}
