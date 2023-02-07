package coupon

import (
	"final/memorydb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getCoupon func() []memorydb.Coupon

func (fn getCoupon) Get() []memorydb.Coupon {
	return fn()
}

func GetCouponHandler(gc getCoupon) echo.HandlerFunc {
	return func(c echo.Context) error {
		coupons := gc.Get()
		return c.JSON(http.StatusOK, coupons)
	}
}

type getCouponByID func(int) (*memorydb.Coupon, error)

func (fn getCouponByID) Get(id int) (*memorydb.Coupon, error) {
	return fn(id)
}

func GetCouponByIDHandler(gc getCouponByID) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		coupon, err := gc.Get(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, coupon)
	}
}
