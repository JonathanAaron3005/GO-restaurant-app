package main

import (
	"github.com/JonathanAaron3005/go-restaurant-app/internal/database"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/JonathanAaron3005/go-restaurant-app/internal/repository/menu"
	orRepo "github.com/JonathanAaron3005/go-restaurant-app/internal/repository/order"
	rUsecase "github.com/JonathanAaron3005/go-restaurant-app/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)
	orderRepo := orRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo)

	handler := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":14045"))
}

const (
	dbAddress = "host=localhost port=5432 user=postgres password=ja3005 dbname=go_resto_app sslmode=disable"
)
