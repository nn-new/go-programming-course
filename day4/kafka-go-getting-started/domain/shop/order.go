package shop

import (
	"time"

	"github.com/google/uuid"
)

// Example Send message with JSON
// {
// 	"ordertime": 1497014222380,
// 	"orderid": 18,
// 	"itemid": "Item_184",
// 	"address": {
// 		"city": "Mountain View",
// 		"state": "CA",
// 		"zipcode": 94041
// 	}
// }

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int64  `json:"zipcode"`
}

type Order struct {
	OrderID   uuid.UUID `json:"orderId"`
	ItemID    string    `json:"itemId"`
	OrderTime time.Time `json:"orderTime"`
	Address   Address   `json:"address"`
}

var MockOrder = Order{
	OrderID:   uuid.New(),
	ItemID:    "Item_184",
	OrderTime: time.Now(),
	Address: Address{
		City:    "Mountain View",
		State:   "CA",
		Zipcode: 94041,
	},
}
