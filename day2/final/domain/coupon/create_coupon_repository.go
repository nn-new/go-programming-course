package coupon

import "final/memorydb"

func CreateCoupon(coupon Coupon) int {
	memorydb.Lock.Lock()
	defer memorydb.Lock.Unlock()

	couponID := memorydb.Seq

	memorydb.Coupons[couponID] = &memorydb.Coupon{
		ID:   memorydb.Seq,
		Name: coupon.Name,
	}

	memorydb.Seq++

	return couponID
}
