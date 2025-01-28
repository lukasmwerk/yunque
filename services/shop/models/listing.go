package models

import "github.com/lukasmwerk/yunque/libraries/types"

// Listing is the listing data for a product used to display a listing.
// A listing can contain multiple products under options. The intended use
// is to enable selecting individual variant products and modifying the displayed listing.
type Listing struct {
	// Name is the name of the parent product.
	Name string

	// PriceRange is a range of prices a product can be based on options.
	PriceRange types.PriceRange

	// Variants contains variants of the item.
	Variants ListingVariants

	// ProductDetails contains details of the current product being displayed.
	ProductDetails ProductDetails
}

// Variants are variations to a product for example shoe sizes.
// Note: Foregoing polymorphism-for-the-sake-of-polymorphism on this,
// the category count is limited and each is handled very differently.
// Named structs for ease of testing and longevity.
type ListingVariants struct {
	ProductColors        []ProductColor
	ProductManufacturers []ProductManufacturer
	ProductQualities     []ProductQuality
	ProductSizes         []ProductSize
	ProductTypes         []ProductType
}
