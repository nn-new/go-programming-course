package coupon

import (
	"final/memorydb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type updateCouponByID func(Coupon) (*memorydb.Coupon, error)

func (fn updateCouponByID) Update(coupon Coupon) (*memorydb.Coupon, error) {
	return fn(coupon)
}

func UpdateCouponByIDHandler(uc updateCouponByID) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var coupon Coupon
		if err := c.Bind(&coupon); err != nil {
			return err
		}
		coupon.ID = id

		cRes, err := uc.Update(coupon)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, cRes)
	}
}
