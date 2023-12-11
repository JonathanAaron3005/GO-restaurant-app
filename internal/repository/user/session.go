package user

import (
	"time"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (user *userRepo) CreateUserSession(userID string) (model.UserSession, error) {
	accessToken, err := user.generateAccessToken(userID)

	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (user *userRepo) generateAccessToken(userID string) (string, error) {
	accessTokenExp := time.Now().Add(user.accessExp).Unix()

	accessClaims := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExp,
		},
	}

	accessJWT := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)

	return accessJWT.SignedString(user.signKey)
}
