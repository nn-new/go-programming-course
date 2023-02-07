package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Coupon struct {
	ID int64 `param:"id" query:"id" form:"id" json:"id"`
}

func main() {
	e := echo.New()

	e.GET("/coupons/:id", func(c echo.Context) error {
		userID := c.Request().Header.Get("USER_ID")

		ID := c.Param("id")
		return c.String(http.StatusOK, fmt.Sprintf("/coupons/%s with header %s", ID, userID))
	})

	e.POST("/coupons/:id", createCouponsHandler)
	e.POST("/upload", uploadHandler)
	e.POST("/uploads", uploadsHandler)

	g := e.Group("/admin")

	g.GET("/coupons/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/admin/coupons/:id")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func createCouponsHandler(c echo.Context) error {
	var coupon Coupon
	err := c.Bind(&coupon)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Get Coupon: %d", coupon.ID))
}

func uploadHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	// Do something...

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
}

func uploadsHandler(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Do something...
		fmt.Println(file.Filename)
	}

	return c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files.", len(files)))
}
