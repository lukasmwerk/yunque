package handlers

import (
	"context"
	"errors"
	"github.com/lukasmwerk/yunque/libraries/types"
	"github.com/lukasmwerk/yunque/services/shop/stores"
)

// CacheProductInCart preemptively inserts a product into the user's cart with a frequency of zero.
// The inserted product is stored in a map of the cart, and if a user chooses to add an amount of
// the item to their cart, the amount for key productID is simply updated.
func CacheProductInCart(ctx context.Context, db *stores.Session, userID types.ID, productID types.ID) error {
	return stores.CacheProductInCart(ctx, db, userID, productID)
}

// RemoveProductFromCart deletes a product from the user cart, both in the cache and persistent store.
func RemoveProductFromCart(ctx context.Context, db *stores.Session, userID types.ID, productID types.ID) error {
	return stores.RemoveProductFromCart(ctx, db, userID, productID)
}

// SetProductAmountInCart updates the user's cart with an amount of the passed product.
func SetProductAmountInCart(ctx context.Context, db *stores.Session, userID types.ID, productID types.ID, amount int) error {
	defer SnapshotCart(ctx, db, userID) // should fail quietly
	return stores.SetProductAmountInCart(ctx, db, userID, productID, amount)
}

// SnapshotCart saves a snapshot of the user's cart (over the network) to the persistent store.
func SnapshotCart(ctx context.Context, db *stores.Session, userID types.ID) error {
	return stores.SnapshotCart(ctx, db, userID)
}

// ClearUserCart removes all products from a user's cart. No undo.
func ClearUserCart(ctx context.Context, db *stores.Session, userID types.ID) error {
	return stores.ClearUserCart(ctx, db, userID)
}

// ViewCart retrieves the cart for a user.
func ViewCart(ctx context.Context, db *stores.Session, userID types.ID) error {
	cart, err := stores.GetCart(ctx, db, userID) // TODO do something with this cart
	if err != nil {
		return err
	} else if cart == nil {
		return errors.New("") // TODO add error
	}
	return nil
}
