package coupon

import (
	"net/http"

	customValidate "final/validator"

	"github.com/labstack/echo/v4"
)

type createCoupon func(Coupon) int

func (fn createCoupon) Create(coupon Coupon) int {
	return fn(coupon)
}

func CreateCouponHandler(gc createCoupon) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon Coupon
		if err := c.Bind(&coupon); err != nil {
			return err
		}

		err := c.Validate(coupon)
		mapErrs := customValidate.MapErrorMessage(err)
		if len(mapErrs) > 0 {
			return c.JSON(http.StatusBadRequest, mapErrs)
		}

		coupons := gc.Create(coupon)
		return c.JSON(http.StatusOK, coupons)
	}
}
