package user

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var request LoginRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if request.UserName != "jon" || request.Password != "shhh!" {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaims{
		"Jon Snow",
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Token: t,
	})
}
