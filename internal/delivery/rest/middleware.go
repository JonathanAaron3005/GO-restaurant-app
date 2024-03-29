package rest

import (
	"context"
	"net/http"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model/constant"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/tracing"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/usecase/resto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://retoku.com"},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))
}

type authMiddleware struct {
	restoUsecase resto.Usecase
}

func GetAuthMiddleware(restoUseCase resto.Usecase) *authMiddleware {
	return &authMiddleware{
		restoUsecase: restoUseCase,
	}
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.CreateSpan(c.Request().Context(), "CheckAuth")
		defer span.End()

		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		userID, err := am.restoUsecase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID) //membuat context baru dengan userID sbg value
		c.SetRequest(c.Request().WithContext(authContext))                                       //set context lama dengan context baru yaitu authContext

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}
