package models

import (
	"github.com/lukasmwerk/yunque/libraries/types"
	"time"
)

type Cart struct {
	// UserID is a unique identifier of the user that owns this cart.
	UserID types.ID

	// Items are products realized in a cart.
	Items []struct {
		Amount       int // Amount of this item in the cart
		ProductID    types.ID
		ProductName  string      // (optimization)
		ProductPrice types.Price // (optimization)
	}

	// Subtotal is the summed price of all the items in the cart,
	// and does not include deals, coupons or fees.
	Subtotal int

	// LastSnap is the last time this cart was snapshot to the persistent store
	LastSnap time.Time
}
