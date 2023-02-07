package memorydb

import "sync"

type Coupon struct {
	ID   int
	Name string
}

type DB struct {
	Coupon Coupon
}

var (
	Coupons = map[int]*Coupon{}
	Seq     = 1
	Lock    = sync.Mutex{}
)
