package user

import (
	"context"
	"errors"
	"time"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/tracing"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (user *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateUserSession")
	defer span.End()

	accessToken, err := user.generateAccessToken(ctx, userID)

	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (user *userRepo) generateAccessToken(ctx context.Context, userID string) (string, error) {
	_, span := tracing.CreateSpan(ctx, "generateAccessToken")
	defer span.End()

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

func (user *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	_, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

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
