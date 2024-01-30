package user

import (
	"errors"
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

func (user *userRepo) CheckSession(data model.UserSession) (userID string, err error) {
	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return &user.signKey.PublicKey, nil
	})

	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	//cek apakah dia bisa dijadiin Claims (claims ny ada ato tidak)
	if !ok {
		return "", errors.New("unauthorized")
	}

	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}

	return "", errors.New("unauthorized")
}
