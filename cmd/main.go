package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/database"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/rahmathidayat1203/go-restaurant-app/internal/repository/menu"
	rUsecase "github.com/rahmathidayat1203/go-restaurant-app/internal/usecase/resto"
)

const (
	dsn = "host=localhost port=5433 user=postgres password=031299 dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dsn)
	menuRepo := mRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUseCase(menuRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
