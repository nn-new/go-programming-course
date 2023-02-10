package user

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type getUserFunc func(context.Context, string) (User, error)

func (fn getUserFunc) GetUser(ctx context.Context, userName string) (User, error) {
	return fn(ctx, userName)
}

func Login(svc getUserFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request LoginRequest
		if err := c.Bind(&request); err != nil {
			return err
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := svc.GetUser(c.Request().Context(), request.UserName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		password, _ := base64.StdEncoding.DecodeString(user.Password)
		if string(password) != request.Password {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		isAdmin := user.Role == Admin

		claims := &JwtCustomClaims{
			user.UserName,
			isAdmin,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, LoginResponse{
			Token: t,
		})
	}
}
