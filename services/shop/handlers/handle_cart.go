package handlers

import (
	"context"
	"github.com/lukasmwerk/yunque/libraries/types"
)

// CacheProductInCart preemptively inserts a product into the user's cart with a frequency of zero.
// The inserted product is stored in a map of the cart, and if a user chooses to add an amount of
// the item to their cart, the amount for key productID is simply updated.
func CacheProductInCart(userID types.ID, productID types.ID, cxt context.Context) error {
	return nil
}

// RemoveProductFromCart deletes a product from the user cart, both in the cache and persistent store.
func RemoveProductFromCart(userID types.ID, productID types.ID, cxt context.Context) error {
	// TODO: delete from cache
	err := SnapshotCart(cxt)
	return err
}

// SetProductAmountInCart updates the user's cart with an amount of the passed product.
func SetProductAmountInCart(userID types.ID, productID types.ID, amount int, cxt context.Context) error {
	defer SnapshotCart(cxt) // should fail quietly
	return nil
}

// SnapshotCart saves a snapshot of the user's cart to the persistent store.
func SnapshotCart(cxt context.Context) error {
	return nil
}

// ClearUserCart removes all products from a user's cart. No undo.
func ClearUserCart(userID types.ID, cxt context.Context) error {
	return nil
}

// ViewCart retrieves the cart for a user.
func ViewCart(userID types.ID, cxt context.Context) error {
	return nil
}
