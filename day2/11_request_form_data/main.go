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

	e.POST("/coupons", createCouponsHandler)

	g := e.Group("/admin")

	g.GET("/coupons/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/admin/coupons/:id")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func createCouponsHandler(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}
