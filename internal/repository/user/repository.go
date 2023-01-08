package user

import "github.com/rahmathidayat1203/go-restaurant-app/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegister(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
	VerifyingUser(username,password string,userData model.User)(bool,error)
	GetUserData(username string)(model.User,error)
	CreateUserSession(userID string)(model.UserSession ,error)
}
