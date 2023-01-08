package user

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rahmathidayat1203/go-restaurant-app/internal/model"
)

type Claims struct {
	jwt.StandardClaims
}

func (ur *userRepo) CreateUserSession(userID string) (model.UserSession, error) {
	accessToken, err := ur.generateAccessToken(userID)

	if err != nil {
		return model.UserSession{}, nil
	}

	return model.UserSession{
		JwtToken: accessToken,
	}, nil
}

func (ur *userRepo) generateAccessToken(userID string) (string, error) {
	accessTokenExp := time.Now().Add(ur.accessExp).Unix()
	accessClaims := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExp,
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)

	return accessJwt.SignedString(ur.signKey)

}

func (ur *userRepo) CheckSession(data model.UserSession) (userID string, err error) {
	accessToken, err := jwt.ParseWithClaims(data.JwtToken, &Claims{})
}
