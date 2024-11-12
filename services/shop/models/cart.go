package models

import (
	"github.com/lukasmwerk/yunque/libraries/types"
	"time"
)

// Item is a product representation in a cart.
// It holds the amount, and a limited subset of product fields.
type Item struct {
	// Amount of this item in the cart.
	Amount int

	// ProductID of the product that this item is.
	ProductID types.ID

	ProductName  string      // (for optimization)
	ProductPrice types.Price // (for optimization)
}

type Cart struct {
	// UserID is a unique identifier of the user that owns this cart.
	UserID types.ID

	// All the items currently in the Cart.
	Items []Item

	// Subtotal is the summed price of all the items in the cart,
	// and does not include deals, coupons or fees.
	Subtotal int

	// LastSnap is the last time this cart was snapshot to the persistent store.
	LastSnap time.Time
}
