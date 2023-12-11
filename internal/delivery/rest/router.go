package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenuList)
	menuGroup.POST("", handler.AddNewMenu)

	orderGroup := e.Group("/order")
	orderGroup.GET("", handler.getAllOrdersInfo)
	orderGroup.GET("/:orderID", handler.GetOrderInfo)
	orderGroup.POST("", handler.Order)

	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}
