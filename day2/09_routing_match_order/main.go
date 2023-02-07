package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/coupons/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/coupons/:id")
	})

	e.GET("/coupons/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/coupons/new")
	})

	e.GET("/coupons/1/files/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/coupons/1/files/*")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
