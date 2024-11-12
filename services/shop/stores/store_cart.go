package stores

import (
	"context"
	"github.com/lukasmwerk/yunque/libraries/types"
	"github.com/lukasmwerk/yunque/services/shop/models"
)

// Database implementation for carts.

func CacheProductInCart(ctx context.Context, db *Session, userID types.ID, productID types.ID) error {
	return nil
}

func RemoveProductFromCart(ctx context.Context, db *Session, userID types.ID, productID types.ID) error {
	return nil
}

func SetProductAmountInCart(ctx context.Context, db *Session, userID types.ID, productID types.ID, amount int) error {
	return nil
}

func SnapshotCart(ctx context.Context, db *Session, userID types.ID) error {
	return nil
}

func ClearUserCart(ctx context.Context, db *Session, id types.ID) error {
	return nil
}

func GetCart(ctx context.Context, db *Session, id types.ID) (*models.Cart, error) {
	return nil, nil
}
