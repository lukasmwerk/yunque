package handlers

import (
	"context"
	"github.com/lukasmwerk/yunque/libraries/types"
	"github.com/lukasmwerk/yunque/services/shop/stores"
)

// GetProductListing returns the listing data for a product.
func GetProductListing(ctx context.Context, db *stores.Session, productID types.ID) error {
	return stores.GetProductListing(ctx, db, productID)
}
