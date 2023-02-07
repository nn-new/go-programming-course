package coupon

import (
	"final/memorydb"
)

func DeleteCoupon(id int) {
	memorydb.Lock.Lock()
	defer memorydb.Lock.Unlock()

	delete(memorydb.Coupons, id)
}
