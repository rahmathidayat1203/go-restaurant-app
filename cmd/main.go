package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/database"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/rahmathidayat1203/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/rahmathidayat1203/go-restaurant-app/internal/repository/order"
	uRepo "github.com/rahmathidayat1203/go-restaurant-app/internal/repository/user"
	rUsecase "github.com/rahmathidayat1203/go-restaurant-app/internal/usecase/resto"
)

const (
	dsn = "host=localhost port=5433 user=postgres password=031299 dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	db := database.GetDB(dsn)
	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}
	restoUsecase := rUsecase.GetUseCase(menuRepo, orderRepo, userRepo)
	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddleware(e)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
