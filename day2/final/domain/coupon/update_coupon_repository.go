package coupon

import (
	"errors"
	"final/memorydb"
)

func UpdateCoupon(coupon Coupon) (*memorydb.Coupon, error) {
	memorydb.Lock.Lock()
	defer memorydb.Lock.Unlock()

	oldCoupon := memorydb.Coupons[coupon.ID]
	if oldCoupon == nil {
		return nil, errors.New("coupon not found")
	}

	memorydb.Coupons[coupon.ID] = &memorydb.Coupon{
		ID:   coupon.ID,
		Name: coupon.Name,
	}

	return memorydb.Coupons[coupon.ID], nil
}
