package stores

import (
	"context"
	"github.com/lukasmwerk/yunque/libraries/types"
	"github.com/lukasmwerk/yunque/services/shop/models"
)

func GetProductListing(ctx context.Context, db *Session, id types.ID) (*models.Listing, error) {
	return nil, nil
}

func GetProduct(ctx context.Context, db *Session, id types.ID) (*models.Product, error) {
	return nil, nil
}
