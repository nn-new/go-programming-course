package main

import (
	"final/domain/coupon"
	"final/domain/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Login route
	e.POST("/login", user.Login)

	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/coupons", coupon.GetCouponHandler(coupon.GetCoupon))
	e.POST("/coupons", coupon.CreateCouponHandler(coupon.CreateCoupon))
	e.GET("/coupons/:id", coupon.GetCouponByIDHandler(coupon.GetCouponByID))
	e.PUT("/coupons/:id", coupon.UpdateCouponByIDHandler(coupon.UpdateCoupon))
	e.DELETE("/coupons/:id", coupon.DeleteCouponHandler(coupon.DeleteCoupon))

	// Restricted group
	r := e.Group("/restricted")
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user.JwtCustomClaims)
		},
		SigningMethod: jwt.SigningMethodHS512.Alg(),
		SigningKey:    []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*user.JwtCustomClaims)
		name := claims.Name
		return c.String(http.StatusOK, "Welcome "+name+"!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
