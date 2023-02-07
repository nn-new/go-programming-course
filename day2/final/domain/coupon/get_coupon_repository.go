package coupon

import (
	"errors"
	"final/memorydb"
)

func GetCoupon() []memorydb.Coupon {
	memorydb.Lock.Lock()
	defer memorydb.Lock.Unlock()

	flat := []memorydb.Coupon{}
	for _, value := range memorydb.Coupons {
		flat = append(flat, *value)
	}
	return flat
}

func GetCouponByID(id int) (*memorydb.Coupon, error) {
	memorydb.Lock.Lock()
	defer memorydb.Lock.Unlock()

	coupon := memorydb.Coupons[id]
	if coupon == nil {
		return nil, errors.New("coupon not found")
	}
	return coupon, nil
}
