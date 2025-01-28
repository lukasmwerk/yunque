package handlers

import (
	"context"
	"errors"
	"github.com/lukasmwerk/yunque/libraries/types"
	"github.com/lukasmwerk/yunque/services/shop/models"
	"github.com/lukasmwerk/yunque/services/shop/stores"
)

// GetProductListing returns the listing data for a product.
func GetProductListing(ctx context.Context, db *stores.Session, productID types.ID) (*models.Listing, error) {
	return stores.GetProductListing(ctx, db, productID)
}

// GetProduct returns a product and its info
func GetProduct(ctx context.Context, db *stores.Session, productID types.ID) (*models.Product, error) {
	return stores.GetProduct(ctx, db, productID)
}

// BuildProductListing Builds a listing from passed products, with the first product as the default product for the listing
func BuildProductListing(ctx context.Context, db *stores.Session, productIDs []types.ID) (*models.Listing, error) {
	if len(productIDs) == 0 {
		return nil, errors.New("no product ids provided")
	}

	listing := &models.Listing{}

	for _, productID := range productIDs {
		product, err := stores.GetProduct(ctx, db, productID)
		if err != nil {
			return nil, err
		}
		listing.PriceRange.Include(product.Price)
		addVariants(listing, product)
	}

	return listing, nil
}

func AddListing(ctx context.Context, db *stores.Session, listing *models.Listing) error {
	return nil
}

// addVariants is a helper that adds product tags to a listing's variants.
func addVariants(listing *models.Listing, product *models.Product) {
	listing.Variants.ProductColors = append(listing.Variants.ProductColors, product.Tags.Color)
	listing.Variants.ProductManufacturers = append(listing.Variants.ProductManufacturers, product.Tags.Manufacturer)
	listing.Variants.ProductQualities = append(listing.Variants.ProductQualities, product.Tags.Quality)
	listing.Variants.ProductSizes = append(listing.Variants.ProductSizes, product.Tags.Size)
	listing.Variants.ProductTypes = append(listing.Variants.ProductTypes, product.Tags.Type)
}

func AddProductsToListing(ctx context.Context, db *stores.Session, productIDs []types.ID, listing *models.Listing) error {
	return nil
}
