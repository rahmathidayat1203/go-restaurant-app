package user

import "github.com/rahmathidayat1203/go-restaurant-app/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegister(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
}
