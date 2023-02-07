package coupon

type Coupon struct {
	ID     int    `json:"id"`
	Name   string `json:"name" validate:"required"`
	Amount int    `json:"amount" validate:"gte=20,lte=100"`
}
