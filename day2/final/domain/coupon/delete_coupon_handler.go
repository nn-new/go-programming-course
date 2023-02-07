package coupon

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type deleteCouponByID func(int)

func (fn deleteCouponByID) Delete(id int) {
	fn(id)
}

func DeleteCouponHandler(dc deleteCouponByID) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		dc.Delete(id)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "OK",
		})
	}
}
