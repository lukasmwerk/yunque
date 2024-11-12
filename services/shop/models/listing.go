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

	// Options are variations to a product for example shoe sizes.
	// Note: Foregoing polymorphism for the sake of polymorphism on this,
	// the category count is limited and each is handled very differently.
	// Named structs for ease of testing and longevity.
	Options struct {
		ProductColors        []ProductColor
		ProductManufacturers []ProductManufacturer
		ProductQualities     []ProductQuality
		ProductSizes         []ProductSize
		ProductTypes         []ProductType
	}

	// ProductDetails contains details of the current product being displayed.
	ProductDetails ProductDetails
}

// ProductDetails contains variant-specific data. For example, a product of a certain
// color may have its own price or description, but we still want to display the parent
// listing that contains the same options lists.
type ProductDetails struct {
	ProductID types.ID
	Price     types.Price
}

type ProductColor struct {
	ID             types.ID // local value, since a color is not part of a predefined set
	Name           string
	Thumbnail      types.Thumbnail
	ProductDetails ProductDetails
}

type ProductManufacturer struct {
	ManufacturerID types.ID
	Name           string
	ProductDetails ProductDetails
}

type ProductQuality struct {
	ID             types.ID // local value
	Name           string
	ProductDetails ProductDetails
}

type ProductSize struct {
	SizeID         types.ID
	Name           string // TODO: add size model or type
	ProductDetails ProductDetails
}

type ProductType struct {
	ID             types.ID
	Name           string
	ProductDetails ProductDetails
}
